package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateModuleRolePermission(id string, req dto.UpdateModuleRolePermissionRequest, updatedBy uuid.UUID) (*dto.ModuleRolePermissionDTO, error) {
	db := config.DB

	permID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var permission models.ModuleRolePermission
	if err := db.Where("id = ? AND is_deleted = FALSE", permID).First(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("permiso no encontrado")
		}
		return nil, err
	}

	if req.PermissionType != nil {
		permission.PermissionType = *req.PermissionType
	}

	if err := db.Save(&permission).Error; err != nil {
		return nil, err
	}

	return GetModuleRolePermissionByID(id)
}