package services

import (
	"server/internal/config"
	"server/internal/models"
)

type OrganicUnitRow struct {
	models.OrganicUnit
	UsersCount int64 `gorm:"column:users_count"`
}

func GetOrganicUnits(page, pageSize int, isDeleted bool) ([]OrganicUnitRow, int64, error) {
	db := config.DB

	var total int64
	base := db.Model(&models.OrganicUnit{}).Where("is_deleted = ?", isDeleted)
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var rows []OrganicUnitRow
	q := db.
		Table("organic_units ou").
		Select(`
			ou.*,
			COUNT(u.id) AS users_count
		`).
		Joins(`
			LEFT JOIN users u
				ON u.organic_unit_id = ou.id
				AND u.is_deleted = FALSE
		`).
		Where("ou.is_deleted = ?", isDeleted).
		Group("ou.id").
		Order("ou.created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := q.Scan(&rows).Error; err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}
