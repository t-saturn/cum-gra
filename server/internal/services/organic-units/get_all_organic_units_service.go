package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"strconv"
)

func GetAllOrganicUnits(onlyActive bool) ([]dto.OrganicUnitSelectDTO, error) {
	db := config.DB

	query := db.Model(&models.OrganicUnit{}).Where("is_deleted = ?", false)

	if onlyActive {
		query = query.Where("is_active = ?", true)
	}

	var units []models.OrganicUnit
	if err := query.Order("name ASC").Find(&units).Error; err != nil {
		return nil, err
	}

	result := make([]dto.OrganicUnitSelectDTO, 0, len(units))
	for _, u := range units {
		var parentID *string
		if u.ParentID != nil {
			pid := strconv.FormatUint(uint64(*u.ParentID), 10)
			parentID = &pid
		}

		result = append(result, dto.OrganicUnitSelectDTO{
			ID:       strconv.FormatUint(uint64(u.ID), 10),
			Name:     u.Name,
			Acronym:  u.Acronym,
			ParentID: parentID,
			IsActive: u.IsActive,
		})
	}

	return result, nil
}