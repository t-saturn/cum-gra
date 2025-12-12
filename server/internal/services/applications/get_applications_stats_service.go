package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetApplicationsStats() (*dto.ApplicationStatsResponse, error) {
	db := config.DB

	var totalApps int64
	if err := db.Model(&models.Application{}).
		Count(&totalApps).Error; err != nil {
		return nil, err
	}

	var activeApps int64
	if err := db.Model(&models.Application{}).
		Where("is_deleted = ? AND status = ?", false, "active").
		Count(&activeApps).Error; err != nil {
		return nil, err
	}

	var inactiveApps int64
	if err := db.Model(&models.Application{}).
		Where("is_deleted = ? AND status = ?", false, "inactive").
		Count(&inactiveApps).Error; err != nil {
		return nil, err
	}

	var deletedApps int64
	if err := db.Model(&models.Application{}).
		Where("is_deleted = ?", true).
		Count(&deletedApps).Error; err != nil {
		return nil, err
	}

	return &dto.ApplicationStatsResponse{
		TotalApplications:    totalApps,
		ActiveApplications:   activeApps,
		InactiveApplications: inactiveApps,
		DeletedApplications:  deletedApps,
	}, nil
}