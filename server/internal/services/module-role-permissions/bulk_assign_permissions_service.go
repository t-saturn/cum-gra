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

func BulkAssignPermissions(req dto.BulkAssignPermissionsRequest, createdBy uuid.UUID) (*dto.BulkAssignPermissionsResponse, error) {
	db := config.DB

	roleID, err := uuid.Parse(req.ApplicationRoleID)
	if err != nil {
		return nil, errors.New("application_role_id inválido")
	}

	// Verificar que el rol existe
	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", roleID).First(&role).Error; err != nil {
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

	response := &dto.BulkAssignPermissionsResponse{
		Created: 0,
		Skipped: 0,
		Failed:  0,
		Details: []dto.ModuleRolePermissionDTO{},
	}

	for _, moduleIDStr := range req.ModuleIDs {
		moduleID, err := uuid.Parse(moduleIDStr)
		if err != nil {
			response.Failed++
			continue
		}

		// Verificar que el módulo existe
		var module models.Module
		if err := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; err != nil {
			response.Failed++
			continue
		}

		// Verificar que el módulo pertenece a la misma aplicación
		if module.ApplicationID == nil || *module.ApplicationID != role.ApplicationID {
			response.Failed++
			continue
		}

		// Verificar si ya existe
		var exists int64
		if err := db.Model(&models.ModuleRolePermission{}).
			Where("module_id = ? AND application_role_id = ? AND is_deleted = FALSE", moduleID, roleID).
			Count(&exists).Error; err != nil {
			response.Failed++
			continue
		}

		if exists > 0 {
			response.Skipped++
			continue
		}

		// Crear el permiso
		permission := models.ModuleRolePermission{
			ID:                uuid.New(),
			ModuleID:          moduleID,
			ApplicationRoleID: roleID,
			PermissionType:    req.PermissionType,
			CreatedAt:         time.Now(),
			IsDeleted:         false,
		}

		if err := db.Create(&permission).Error; err != nil {
			response.Failed++
			continue
		}

		response.Created++
		response.Details = append(response.Details, 
			mapper.ToModuleRolePermissionDTO(permission, &module, &role, &app))
	}

	return response, nil
}