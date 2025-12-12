package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetModulesStats() (*dto.ModulesStatsResponse, error) {
	db := config.DB

	var totalModules int64
	if err := db.Model(&models.Module{}).
		Count(&totalModules).Error; err != nil {
		return nil, err
	}

	var activeModules int64
	if err := db.Model(&models.Module{}).
		Where("deleted_at IS NULL AND status = ?", "active").
		Count(&activeModules).Error; err != nil {
		return nil, err
	}

	var deletedModules int64
	if err := db.Model(&models.Module{}).
		Where("deleted_at IS NOT NULL").
		Count(&deletedModules).Error; err != nil {
		return nil, err
	}

	var totalUsers int64
	if err := db.Table("module_role_permissions mrp").
		Joins("JOIN user_application_roles uar ON uar.application_role_id = mrp.application_role_id AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL").
		Where("mrp.is_deleted = FALSE").
		Distinct("uar.user_id").
		Count(&totalUsers).Error; err != nil {
		return nil, err
	}

	return &dto.ModulesStatsResponse{
		TotalModules:   totalModules,
		ActiveModules:  activeModules,
		DeletedModules: deletedModules,
		TotalUsers:     totalUsers,
	}, nil
}