package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetUbigeosStats() (*dto.UbigeoStatsResponse, error) {
	db := config.DB

	var totalUbigeos int64
	if err := db.Model(&models.Ubigeo{}).
		Count(&totalUbigeos).Error; err != nil {
		return nil, err
	}

	var totalDepartments int64
	if err := db.Model(&models.Ubigeo{}).
		Distinct("department").
		Count(&totalDepartments).Error; err != nil {
		return nil, err
	}

	var totalProvinces int64
	if err := db.Model(&models.Ubigeo{}).
		Distinct("department", "province").
		Count(&totalProvinces).Error; err != nil {
		return nil, err
	}

	var totalDistricts int64
	if err := db.Model(&models.Ubigeo{}).
		Distinct("department", "province", "district").
		Count(&totalDistricts).Error; err != nil {
		return nil, err
	}

	return &dto.UbigeoStatsResponse{
		TotalUbigeos:     totalUbigeos,
		TotalDepartments: totalDepartments,
		TotalProvinces:   totalProvinces,
		TotalDistricts:   totalDistricts,
	}, nil
}