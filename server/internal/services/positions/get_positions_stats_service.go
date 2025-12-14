package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetStructuralPositionsStats() (*dto.StructuralPositionsStatsResponse, error) {
	db := config.DB

	var totalPositions int64
	if err := db.Model(&models.StructuralPosition{}).
		Count(&totalPositions).Error; err != nil {
		return nil, err
	}

	var activePositions int64
	if err := db.Model(&models.StructuralPosition{}).
		Where("is_deleted = FALSE AND is_active = TRUE").
		Count(&activePositions).Error; err != nil {
		return nil, err
	}

	var deletedPositions int64
	if err := db.Model(&models.StructuralPosition{}).
		Where("is_deleted = TRUE").
		Count(&deletedPositions).Error; err != nil {
		return nil, err
	}

	var assignedEmployees int64
	if err := db.Table("user_details ud").
		Joins("JOIN structural_positions sp ON sp.id = ud.structural_position_id").
		Where("sp.is_deleted = FALSE").
		Distinct("ud.user_id").
		Count(&assignedEmployees).Error; err != nil {
		return nil, err
	}

	return &dto.StructuralPositionsStatsResponse{
		TotalPositions:    totalPositions,
		ActivePositions:   activePositions,
		DeletedPositions:  deletedPositions,
		AssignedEmployees: assignedEmployees,
	}, nil
}