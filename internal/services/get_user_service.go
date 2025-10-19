package services

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/models"
)

func GetUsers(page, pageSize int, isDeleted bool) ([]models.User, int64, error) {
	db := config.DB

	var total int64
	base := db.Model(&models.User{}).Where("is_deleted = ?", isDeleted)
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var users []models.User
	q := db.
		Preload("OrganicUnit").
		Preload("StructuralPosition").
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := q.Find(&users).Error; err != nil {
		return nil, total, err
	}

	return users, total, nil
}
