package services

import (
	"time"

	"server/internal/config"
	"server/internal/models"
)

func GetUserStats() (total, active, suspended, newLastMonth int64, err error) {
	db := config.DB
	since := time.Now().AddDate(0, -1, 0)

	if err = db.Model(&models.User{}).
		Where("is_deleted = ?", false).
		Count(&total).Error; err != nil {
		return
	}

	if err = db.Model(&models.User{}).
		Where("is_deleted = ? AND status = ?", false, "active").
		Count(&active).Error; err != nil {
		return
	}

	if err = db.Model(&models.User{}).
		Where("is_deleted = ? AND status = ?", false, "inactive").
		Count(&suspended).Error; err != nil {
		return
	}

	if err = db.Model(&models.User{}).
		Where("is_deleted = ? AND created_at >= ?", false, since).
		Count(&newLastMonth).Error; err != nil {
		return
	}

	return
}
