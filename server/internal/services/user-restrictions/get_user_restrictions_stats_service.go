package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetUserModuleRestrictionsStats() (*dto.UserRestrictionsStatsDTO, error) {
	db := config.DB
	now := time.Now()

	var totalRestrictions int64
	if err := db.Model(&models.UserModuleRestriction{}).
		Count(&totalRestrictions).Error; err != nil {
		return nil, err
	}

	var activeRestrictions int64
	if err := db.Model(&models.UserModuleRestriction{}).
		Where("is_deleted = FALSE AND (expires_at IS NULL OR expires_at > ?)", now).
		Count(&activeRestrictions).Error; err != nil {
		return nil, err
	}

	var restrictedUsers int64
	if err := db.Model(&models.UserModuleRestriction{}).
		Where("is_deleted = FALSE AND (expires_at IS NULL OR expires_at > ?)", now).
		Distinct("user_id").
		Count(&restrictedUsers).Error; err != nil {
		return nil, err
	}

	var deletedRestrictions int64
	if err := db.Model(&models.UserModuleRestriction{}).
		Where("is_deleted = TRUE").
		Count(&deletedRestrictions).Error; err != nil {
		return nil, err
	}

	return &dto.UserRestrictionsStatsDTO{
		TotalRestrictions:   totalRestrictions,
		ActiveRestrictions:  activeRestrictions,
		RestrictedUsers:     restrictedUsers,
		DeletedRestrictions: deletedRestrictions,
	}, nil
}