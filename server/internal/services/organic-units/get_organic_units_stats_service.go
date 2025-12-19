package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetOrganicUnitsStats() (*dto.OrganicUnitsStatsResponse, error) {
	db := config.DB

	var totalUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Count(&totalUnits).Error; err != nil {
		return nil, err
	}

	var activeUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = FALSE AND is_active = TRUE").
		Count(&activeUnits).Error; err != nil {
		return nil, err
	}

	var deletedUnits int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("is_deleted = TRUE").
		Count(&deletedUnits).Error; err != nil {
		return nil, err
	}

	var totalEmployees int64
	if err := db.Table("user_details ud").
		Joins("JOIN organic_units ou ON ou.id = ud.organic_unit_id").
		Where("ou.is_deleted = FALSE").
		Distinct("ud.user_id").
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