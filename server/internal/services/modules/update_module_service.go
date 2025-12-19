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

func UpdateModule(id string, req dto.UpdateModuleRequest, updatedBy uuid.UUID) (*dto.ModuleWithAppDTO, error) {
	db := config.DB

	moduleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var module models.Module
	if err := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("módulo no encontrado")
		}
		return nil, err
	}

	// Validar nombre único si se está actualizando
	if req.Name != nil && *req.Name != module.Name {
		var exists int64
		query := db.Model(&models.Module{}).
			Where("name = ? AND id != ? AND deleted_at IS NULL", *req.Name, moduleID)
		
		if module.ApplicationID != nil {
			query = query.Where("application_id = ?", *module.ApplicationID)
		} else {
			query = query.Where("application_id IS NULL")
		}
		
		if err := query.Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe un módulo con este nombre en esta aplicación")
		}
		module.Name = *req.Name
	}

	// Validar parent_id si se está actualizando
	if req.ParentID != nil {
		if *req.ParentID == "" {
			module.ParentID = nil
		} else {
			parsedParentID, err := uuid.Parse(*req.ParentID)
			if err != nil {
				return nil, errors.New("parent_id inválido")
			}
			
			// No puede ser su propio padre
			if parsedParentID == moduleID {
				return nil, errors.New("un módulo no puede ser su propio padre")
			}
			
			var parent models.Module
			if err := db.Where("id = ? AND deleted_at IS NULL", parsedParentID).First(&parent).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("módulo padre no encontrado")
				}
				return nil, err
			}
			module.ParentID = &parsedParentID
		}
	}

	if req.Item != nil {
		module.Item = req.Item
	}
	if req.Route != nil {
		module.Route = req.Route
	}
	if req.Icon != nil {
		module.Icon = req.Icon
	}
	if req.SortOrder != nil {
		module.SortOrder = *req.SortOrder
	}
	if req.Status != nil {
		module.Status = *req.Status
	}

	module.UpdatedAt = time.Now()

	if err := db.Save(&module).Error; err != nil {
		return nil, err
	}

	return GetModuleByID(id)
}