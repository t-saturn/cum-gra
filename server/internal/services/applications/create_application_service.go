package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func CreateApplication(req dto.CreateApplicationRequest, createdBy uuid.UUID) (*dto.ApplicationDTO, error) {
	db := config.DB

	// Verificar si ya existe el client_id
	var exists int64
	if err := db.Model(&models.Application{}).
		Where("client_id = ? AND is_deleted = FALSE", req.ClientID).
		Count(&exists).Error; err != nil {
		return nil, err
	}

	if exists > 0 {
		return nil, errors.New("ya existe una aplicación con este client_id")
	}

	status := "active"
	if req.Status != "" {
		status = req.Status
	}

	app := models.Application{
		ID:           uuid.New(),
		Name:         req.Name,
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret, // TODO: Deberías encriptar esto
		Domain:       req.Domain,
		Logo:         req.Logo,
		Description:  req.Description,
		Status:       status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsDeleted:    false,
	}

	if err := db.Create(&app).Error; err != nil {
		return nil, err
	}

	result := mapper.ToApplicationDTO(app, []dto.AdminUserDTO{}, 0)
	return &result, nil
}