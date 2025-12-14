package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
)

func GetStructuralPositions(page, pageSize int, isDeleted bool, level *int) (*dto.StructuralPositionsListResponse, error) {
	db := config.DB

	query := db.Model(&models.StructuralPosition{}).Where("is_deleted = ?", isDeleted)

	if level != nil {
		query = query.Where("level = ?", *level)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var positions []models.StructuralPosition
	if err := query.
		Order("level ASC, name ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&positions).Error; err != nil {
		return nil, err
	}

	if len(positions) == 0 {
		return &dto.StructuralPositionsListResponse{
			Data:     []dto.StructuralPositionItemDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	positionIDs := make([]uint, 0, len(positions))
	for _, p := range positions {
		positionIDs = append(positionIDs, p.ID)
	}

	// Obtener count de usuarios por posición
	var userCounts []struct {
		PositionID uint  `gorm:"column:structural_position_id"`
		Count      int64 `gorm:"column:count"`
	}
	if err := db.Table("user_details").
		Select("structural_position_id, COUNT(*) as count").
		Where("structural_position_id IN ?", positionIDs).
		Group("structural_position_id").
		Scan(&userCounts).Error; err != nil {
		return nil, err
	}

	userCountMap := make(map[uint]int64)
	for _, uc := range userCounts {
		userCountMap[uc.PositionID] = uc.Count
	}

	rows := make([]models.StructuralPositionRow, 0, len(positions))
	for _, pos := range positions {
		usersCount := userCountMap[pos.ID]
		rows = append(rows, mapper.StructuralPositionToRow(pos, usersCount))
	}

	out := mapper.ToStructuralPositionListDTO(rows)

	return &dto.StructuralPositionsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}