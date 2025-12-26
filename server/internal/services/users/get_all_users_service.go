package services

import (
	"server/internal/config"
	"server/internal/dto"
)

func GetAllUsers(onlyActive bool) ([]dto.UserSelectDTO, error) {
	db := config.DB

	query := db.Table("users").
		Select("users.id, users.email, users.dni, user_details.first_name, user_details.last_name, user_details.status").
		Joins("LEFT JOIN user_details ON users.id = user_details.user_id").
		Where("users.is_deleted = ?", false)

	if onlyActive {
		query = query.Where("user_details.status = ?", "active")
	}

	var results []struct {
		ID        string  `gorm:"column:id"`
		Email     string  `gorm:"column:email"`
		DNI       string  `gorm:"column:dni"`
		FirstName *string `gorm:"column:first_name"`
		LastName  *string `gorm:"column:last_name"`
		Status    string  `gorm:"column:status"`
	}

	if err := query.Order("user_details.first_name ASC, user_details.last_name ASC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	users := make([]dto.UserSelectDTO, 0, len(results))
	for _, r := range results {
		users = append(users, dto.UserSelectDTO{
			ID:        r.ID,
			Email:     r.Email,
			DNI:       r.DNI,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Status:    r.Status,
		})
	}

	return users, nil
}