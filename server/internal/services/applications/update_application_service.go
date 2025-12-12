package srvapplications

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateApplication(id string, req dto.UpdateApplicationRequest, updatedBy uuid.UUID) (*dto.ApplicationDTO, error) {
	db := config.DB

	appID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	// Verificar client_id único si se está actualizando
	if req.ClientID != nil && *req.ClientID != app.ClientID {
		var exists int64
		if err := db.Model(&models.Application{}).
			Where("client_id = ? AND id != ? AND is_deleted = FALSE", *req.ClientID, appID).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe una aplicación con este client_id")
		}
		app.ClientID = *req.ClientID
	}

	if req.Name != nil {
		app.Name = *req.Name
	}
	if req.ClientSecret != nil {
		app.ClientSecret = *req.ClientSecret // TODO: Deberías encriptar esto
	}
	if req.Domain != nil {
		app.Domain = *req.Domain
	}
	if req.Logo != nil {
		app.Logo = req.Logo
	}
	if req.Description != nil {
		app.Description = req.Description
	}
	if req.Status != nil {
		app.Status = *req.Status
	}

	app.UpdatedAt = time.Now()

	if err := db.Save(&app).Error; err != nil {
		return nil, err
	}

	return GetApplicationByID(id)
}