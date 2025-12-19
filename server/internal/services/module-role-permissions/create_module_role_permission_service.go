package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateModuleRolePermission(req dto.CreateModuleRolePermissionRequest, createdBy uuid.UUID) (*dto.ModuleRolePermissionDTO, error) {
	db := config.DB

	moduleID, err := uuid.Parse(req.ModuleID)
	if err != nil {
		return nil, errors.New("module_id inválido")
	}

	roleID, err := uuid.Parse(req.ApplicationRoleID)
	if err != nil {
		return nil, errors.New("application_role_id inválido")
	}

	// Verificar que el módulo existe
	var module models.Module
	if err := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("módulo no encontrado")
		}
		return nil, err
	}

	// Verificar que el rol existe
	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rol no encontrado")
		}
		return nil, err
	}

	// Verificar que el módulo pertenece a la misma aplicación del rol
	if module.ApplicationID == nil || *module.ApplicationID != role.ApplicationID {
		return nil, errors.New("el módulo y el rol deben pertenecer a la misma aplicación")
	}

	// Verificar si ya existe este permiso
	var exists int64
	if err := db.Model(&models.ModuleRolePermission{}).
		Where("module_id = ? AND application_role_id = ? AND is_deleted = FALSE", moduleID, roleID).
		Count(&exists).Error; err != nil {
		return nil, err
	}

	if exists > 0 {
		return nil, errors.New("ya existe un permiso para este módulo y rol")
	}

	permission := models.ModuleRolePermission{
		ID:                uuid.New(),
		ModuleID:          moduleID,
		ApplicationRoleID: roleID,
		PermissionType:    req.PermissionType,
		CreatedAt:         time.Now(),
		IsDeleted:         false,
	}

	if err := db.Create(&permission).Error; err != nil {
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