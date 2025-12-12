package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetRolesAppStats() (*dto.RolesAppStatsResponse, error) {
	db := config.DB
	var totalRoles, activeRoles, adminRoles, assignedUsers int64

	if err := db.Model(&models.ApplicationRole{}).
		Count(&totalRoles).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = FALSE").
		Count(&activeRoles).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = FALSE AND name ILIKE '%admin%'").
		Count(&adminRoles).Error; err != nil {
		return nil, err
	}

	if err := db.Table("user_application_roles uar").
		Joins("JOIN application_roles ar ON ar.id = uar.application_role_id").
		Where("uar.is_deleted = FALSE AND ar.is_deleted = FALSE").
		Distinct("uar.user_id").
		Count(&assignedUsers).Error; err != nil {
		return nil, err
	}

	return &dto.RolesAppStatsResponse{
		TotalRoles:    totalRoles,
		ActiveRoles:   activeRoles,
		AdminRoles:    adminRoles,
		AssignedUsers: assignedUsers,
	}, nil
}
