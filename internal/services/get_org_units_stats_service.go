package services

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
)

func GetOrganicUnitsStats() (*dto.OrganicUnitsStatsResponse, error) {
	db := config.DB
	var totalUnits, activeUnits, deletedUnits, totalEmployees int64

	if err := db.Model(&models.OrganicUnit{}).Count(&totalUnits).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = ? AND is_active = ?", false, true).
		Count(&activeUnits).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = ?", true).
		Count(&deletedUnits).Error; err != nil {
		return nil, err
	}

	if err := db.Table("users u").
		Joins("JOIN organic_units ou ON ou.id = u.organic_unit_id").
		Where("u.is_deleted = FALSE AND ou.is_deleted = FALSE").
		Count(&totalEmployees).Error; err != nil {
		return nil, err
	}

	return &dto.OrganicUnitsStatsResponse{
		TotalOrganicUnits:   totalUnits,
		ActiveOrganicUnits:  activeUnits,
		DeletedOrganicUnits: deletedUnits,
		TotalEmployees:      totalEmployees,
	}, nil
}
