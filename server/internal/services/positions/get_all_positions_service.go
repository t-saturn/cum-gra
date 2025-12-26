package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"strconv"
)

func GetAllPositions(onlyActive bool) ([]dto.PositionSelectDTO, error) {
	db := config.DB

	query := db.Model(&models.StructuralPosition{}).Where("is_deleted = ?", false)

	if onlyActive {
		query = query.Where("is_active = ?", true)
	}

	var positions []models.StructuralPosition
	if err := query.Order("level ASC, name ASC").Find(&positions).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PositionSelectDTO, 0, len(positions))
	for _, p := range positions {
		result = append(result, dto.PositionSelectDTO{
			ID:       strconv.FormatUint(uint64(p.ID), 10),
			Name:     p.Name,
			Code:     p.Code,
			Level:    p.Level,
			IsActive: p.IsActive,
		})
	}

	return result, nil
}