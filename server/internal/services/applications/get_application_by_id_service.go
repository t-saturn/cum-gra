package services

import (
	"errors"
	"strings"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetApplicationByID(id string) (*dto.ApplicationDTO, error) {
	db := config.DB

	appID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var app models.Application
	if err := db.Where("id = ?", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	// Obtener count de usuarios
	var usersCount int64
	if err := db.
		Table("user_application_roles").
		Where("application_id = ? AND is_deleted = FALSE AND revoked_at IS NULL", appID).
		Count(&usersCount).Error; err != nil {
		return nil, err
	}

	// Obtener admins
	var admins []AdminRow
	if err := db.
		Table("application_roles ar").
		Select(`
			ar.application_id,
			uar.user_id,
			u.email,
			u.dni,
			ud.first_name,
			ud.last_name
		`).
		Joins(`JOIN user_application_roles uar ON uar.application_role_id = ar.id AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u ON u.id = uar.user_id AND u.is_deleted = FALSE`).
		Joins(`LEFT JOIN user_details ud ON ud.user_id = u.id`).
		Where("ar.is_deleted = FALSE AND ar.application_id = ? AND ar.name ILIKE ?", appID, "admin%").
		Scan(&admins).Error; err != nil {
		return nil, err
	}

	adminList := make([]dto.AdminUserDTO, 0)
	for _, a := range admins {
		fullName := ""
		if a.FirstName != nil && a.LastName != nil {
			fullName = strings.TrimSpace(*a.FirstName + " " + *a.LastName)
		} else if a.FirstName != nil {
			fullName = *a.FirstName
		} else if a.LastName != nil {
			fullName = *a.LastName
		}
		
		if fullName == "" {
			fullName = a.Email
		}

		adminDTO := dto.AdminUserDTO{
			FullName: fullName,
			DNI:      a.DNI,
			Email:    a.Email,
		}
		adminList = append(adminList, adminDTO)
	}

	result := mapper.ToApplicationDTO(app, adminList, usersCount)
	return &result, nil
}