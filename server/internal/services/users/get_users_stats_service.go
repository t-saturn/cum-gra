package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetUsersStats() (*dto.UsersStatsResponse, error) {
	db := config.DB

	var totalUsers int64
	if err := db.Model(&models.User{}).
		Count(&totalUsers).Error; err != nil {
		return nil, err
	}

	var activeUsers int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = FALSE AND status = ?", "active").
		Count(&activeUsers).Error; err != nil {
		return nil, err
	}

	var suspendedUsers int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = FALSE AND status = ?", "suspended").
		Count(&suspendedUsers).Error; err != nil {
		return nil, err
	}

	// Usuarios nuevos del último mes
	lastMonth := time.Now().AddDate(0, -1, 0)
	var newUsersLastMonth int64
	if err := db.Model(&models.User{}).
		Where("created_at >= ? AND is_deleted = FALSE", lastMonth).
		Count(&newUsersLastMonth).Error; err != nil {
		return nil, err
	}

	return &dto.UsersStatsResponse{
		TotalUsers:        totalUsers,
		ActiveUsers:       activeUsers,
		SuspendedUsers:    suspendedUsers,
		NewUsersLastMonth: newUsersLastMonth,
	}, nil
}