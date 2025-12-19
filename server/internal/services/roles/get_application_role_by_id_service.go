package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetApplicationRoleByID(id string) (*dto.ApplicationRoleDTO, error) {
	db := config.DB

	roleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var role models.ApplicationRole
	if err := db.Where("id = ?", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rol no encontrado")
		}
		return nil, err
	}

	// Obtener application
	var app models.Application
	if err := db.Where("id = ?", role.ApplicationID).First(&app).Error; err != nil {
		return nil, err
	}

	// Obtener count de módulos
	var modulesCount int64
	if err := db.Table("module_role_permissions").
		Where("is_deleted = FALSE AND application_role_id = ?", roleID).
		Count(&modulesCount).Error; err != nil {
		return nil, err
	}

	// Obtener count de usuarios
	var usersCount int64
	if err := db.Table("user_application_roles").
		Where("is_deleted = FALSE AND revoked_at IS NULL AND application_role_id = ?", roleID).
		Count(&usersCount).Error; err != nil {
		return nil, err
	}

	result := mapper.ToApplicationRoleDTO(role, &app, modulesCount, usersCount)
	return &result, nil
}