package services

import (
	"server/internal/config"
	"server/internal/models"
)

func GetPositionsStats() (total, active, deleted, assigned int64, err error) {
	db := config.DB

	if e := db.Model(&models.StructuralPosition{}).Count(&total).Error; e != nil {
		return 0, 0, 0, 0, e
	}

	if e := db.Model(&models.StructuralPosition{}).
		Where("is_deleted = ? AND is_active = ?", false, true).
		Count(&active).Error; e != nil {
		return 0, 0, 0, 0, e
	}

	if e := db.Model(&models.StructuralPosition{}).
		Where("is_deleted = ?", true).
		Count(&deleted).Error; e != nil {
		return 0, 0, 0, 0, e
	}

	if e := db.
		Table("users u").
		Joins("JOIN structural_positions sp ON sp.id = u.structural_position_id").
		Where("u.is_deleted = FALSE AND sp.is_deleted = FALSE").
		Count(&assigned).Error; e != nil {
		return 0, 0, 0, 0, e
	}

	return total, active, deleted, assigned, nil
}
