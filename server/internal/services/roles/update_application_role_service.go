package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateApplicationRole(id string, req dto.UpdateApplicationRoleRequest, updatedBy uuid.UUID) (*dto.ApplicationRoleDTO, error) {
	db := config.DB

	roleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rol no encontrado")
		}
		return nil, err
	}

	// Verificar nombre único si se está actualizando
	if req.Name != nil && *req.Name != role.Name {
		var exists int64
		if err := db.Model(&models.ApplicationRole{}).
			Where("application_id = ? AND name = ? AND id != ? AND is_deleted = FALSE", role.ApplicationID, *req.Name, roleID).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe un rol con este nombre en esta aplicación")
		}
		role.Name = *req.Name
	}

	if req.Description != nil {
		role.Description = req.Description
	}

	role.UpdatedAt = time.Now()

	if err := db.Save(&role).Error; err != nil {
		return nil, err
	}

	return GetApplicationRoleByID(id)
}