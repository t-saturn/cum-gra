package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
	"strconv"
)

func GetOrganicUnits(page, pageSize int, isDeleted bool, parentID *string) (*dto.OrganicUnitsListResponse, error) {
	db := config.DB

	query := db.Model(&models.OrganicUnit{}).Where("is_deleted = ?", isDeleted)

	if parentID != nil && *parentID != "" {
		if *parentID == "null" {
			query = query.Where("parent_id IS NULL")
		} else {
			parentIDUint, err := strconv.ParseUint(*parentID, 10, 32)
			if err != nil {
				return nil, err
			}
			query = query.Where("parent_id = ?", uint(parentIDUint))
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var units []models.OrganicUnit
	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&units).Error; err != nil {
		return nil, err
	}

	if len(units) == 0 {
		return &dto.OrganicUnitsListResponse{
			Data:     []dto.OrganicUnitItemDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	unitIDs := make([]uint, 0, len(units))
	for _, u := range units {
		unitIDs = append(unitIDs, u.ID)
	}

	// Obtener count de usuarios por unidad orgánica
	var userCounts []struct {
		OrganicUnitID uint  `gorm:"column:organic_unit_id"`
		Count         int64 `gorm:"column:count"`
	}
	if err := db.Table("user_details").
		Select("organic_unit_id, COUNT(*) as count").
		Where("organic_unit_id IN ?", unitIDs).
		Group("organic_unit_id").
		Scan(&userCounts).Error; err != nil {
		return nil, err
	}

	userCountMap := make(map[uint]int64)
	for _, uc := range userCounts {
		userCountMap[uc.OrganicUnitID] = uc.Count
	}

	out := make([]dto.OrganicUnitItemDTO, 0, len(units))
	for _, unit := range units {
		usersCount := userCountMap[unit.ID]
		out = append(out, mapper.ToOrganicUnitItemDTO(unit, usersCount))
	}

	return &dto.OrganicUnitsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}