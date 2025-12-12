package services

import (
	"server/internal/config"
	"server/internal/models"
)

func GetPositions(page, pageSize int, isDeleted bool) ([]models.StructuralPositionRow, int64, error) {
	db := config.DB

	var total int64
	if err := db.Model(&models.StructuralPosition{}).
		Where("is_deleted = ?", isDeleted).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var rows []models.StructuralPositionRow
	q := db.
		Table("structural_positions sp").
		Select(`
			sp.*,
			COUNT(u.id) AS users_count
		`).
		Joins(`
			LEFT JOIN users u
				ON u.structural_position_id = sp.id
				AND u.is_deleted = FALSE
		`).
		Where("sp.is_deleted = ?", isDeleted).
		Group("sp.id").
		Order("sp.created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := q.Scan(&rows).Error; err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}
