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

func GetModuleRolePermissionByID(id string) (*dto.ModuleRolePermissionDTO, error) {
	db := config.DB

	permID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var permission models.ModuleRolePermission
	if err := db.Where("id = ?", permID).First(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("permiso no encontrado")
		}
		return nil, err
	}

	// Obtener módulo
	var module models.Module
	if err := db.Where("id = ?", permission.ModuleID).First(&module).Error; err != nil {
		return nil, err
	}

	// Obtener rol
	var role models.ApplicationRole
	if err := db.Where("id = ?", permission.ApplicationRoleID).First(&role).Error; err != nil {
		return nil, err
	}

	// Obtener application
	var app models.Application
	if err := db.Where("id = ?", role.ApplicationID).First(&app).Error; err != nil {
		return nil, err
	}

	result := mapper.ToModuleRolePermissionDTO(permission, &module, &role, &app)
	return &result, nil
}