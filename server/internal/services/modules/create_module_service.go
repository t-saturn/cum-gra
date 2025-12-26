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

func CreateModule(req dto.CreateModuleRequest, createdBy uuid.UUID) (*dto.ModuleWithAppDTO, error) {
	db := config.DB

	// Validar application_id si se proporciona
	var appID *uuid.UUID
	var app *models.Application
	if req.ApplicationID != nil && *req.ApplicationID != "" {
		parsedAppID, err := uuid.Parse(*req.ApplicationID)
		if err != nil {
			return nil, errors.New("application_id inválido")
		}
		
		var appModel models.Application
		if err := db.Where("id = ? AND is_deleted = FALSE", parsedAppID).First(&appModel).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("aplicación no encontrada")
			}
			return nil, err
		}
		appID = &parsedAppID
		app = &appModel
	}

	// Validar parent_id si se proporciona
	var parentID *uuid.UUID
	if req.ParentID != nil && *req.ParentID != "" {
		parsedParentID, err := uuid.Parse(*req.ParentID)
		if err != nil {
			return nil, errors.New("parent_id inválido")
		}
		
		var parent models.Module
		if err := db.Where("id = ? AND deleted_at IS NULL", parsedParentID).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("módulo padre no encontrado")
			}
			return nil, err
		}
		parentID = &parsedParentID
	}

	// Verificar nombre único SOLO para módulos raíz (sin padre) en la misma aplicación
	if parentID == nil {
		var exists int64
		query := db.Model(&models.Module{}).
			Where("name = ? AND deleted_at IS NULL AND parent_id IS NULL", req.Name)
		
		if appID != nil {
			query = query.Where("application_id = ?", *appID)
		} else {
			query = query.Where("application_id IS NULL")
		}
		
		if err := query.Count(&exists).Error; err != nil {
			return nil, err
		}

		if exists > 0 {
			return nil, errors.New("ya existe un módulo raíz con este nombre en esta aplicación")
		}
	}

	status := "active"
	if req.Status != nil {
		status = *req.Status
	}

	sortOrder := 0
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}

	module := models.Module{
		ID:            uuid.New(),
		Item:          req.Item,
		Name:          req.Name,
		Route:         req.Route,
		Icon:          req.Icon,
		ParentID:      parentID,
		ApplicationID: appID,
		SortOrder:     sortOrder,
		Status:        status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := db.Create(&module).Error; err != nil {
		return nil, err
	}

	result := mapper.ToModuleWithAppDTO(module, app, 0)
	return &result, nil
}