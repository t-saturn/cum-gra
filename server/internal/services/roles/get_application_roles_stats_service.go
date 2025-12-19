package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetApplicationRolesStats() (*dto.ApplicationRolesStatsResponse, error) {
	db := config.DB

	var totalRoles int64
	if err := db.Model(&models.ApplicationRole{}).
		Count(&totalRoles).Error; err != nil {
		return nil, err
	}

	var activeRoles int64
	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = FALSE").
		Count(&activeRoles).Error; err != nil {
		return nil, err
	}

	var deletedRoles int64
	if err := db.Model(&models.ApplicationRole{}).
		Where("is_deleted = TRUE").
		Count(&deletedRoles).Error; err != nil {
		return nil, err
	}

	var rolesWithModules int64
	if err := db.Table("application_roles ar").
		Joins("JOIN module_role_permissions mrp ON mrp.application_role_id = ar.id AND mrp.is_deleted = FALSE").
		Where("ar.is_deleted = FALSE").
		Distinct("ar.id").
		Count(&rolesWithModules).Error; err != nil {
		return nil, err
	}

	var rolesWithUsers int64
	if err := db.Table("application_roles ar").
		Joins("JOIN user_application_roles uar ON uar.application_role_id = ar.id AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL").
		Where("ar.is_deleted = FALSE").
		Distinct("ar.id").
		Count(&rolesWithUsers).Error; err != nil {
		return nil, err
	}

	return &dto.ApplicationRolesStatsResponse{
		TotalRoles:       totalRoles,
		ActiveRoles:      activeRoles,
		DeletedRoles:     deletedRoles,
		RolesWithModules: rolesWithModules,
		RolesWithUsers:   rolesWithUsers,
	}, nil
}