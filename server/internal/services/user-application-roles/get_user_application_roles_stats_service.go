package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetUserApplicationRolesStats() (*dto.UserApplicationRolesStatsDTO, error) {
	db := config.DB

	var totalAssignments int64
	if err := db.Model(&models.UserApplicationRole{}).
		Count(&totalAssignments).Error; err != nil {
		return nil, err
	}

	var activeAssignments int64
	if err := db.Model(&models.UserApplicationRole{}).
		Where("is_deleted = FALSE AND revoked_at IS NULL").
		Count(&activeAssignments).Error; err != nil {
		return nil, err
	}

	var revokedAssignments int64
	if err := db.Model(&models.UserApplicationRole{}).
		Where("is_deleted = FALSE AND revoked_at IS NOT NULL").
		Count(&revokedAssignments).Error; err != nil {
		return nil, err
	}

	var deletedAssignments int64
	if err := db.Model(&models.UserApplicationRole{}).
		Where("is_deleted = TRUE").
		Count(&deletedAssignments).Error; err != nil {
		return nil, err
	}

	var usersWithRoles int64
	if err := db.Model(&models.UserApplicationRole{}).
		Where("is_deleted = FALSE AND revoked_at IS NULL").
		Distinct("user_id").
		Count(&usersWithRoles).Error; err != nil {
		return nil, err
	}

	return &dto.UserApplicationRolesStatsDTO{
		TotalAssignments:   totalAssignments,
		ActiveAssignments:  activeAssignments,
		RevokedAssignments: revokedAssignments,
		DeletedAssignments: deletedAssignments,
		UsersWithRoles:     usersWithRoles,
	}, nil
}