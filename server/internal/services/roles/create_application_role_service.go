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

func CreateApplicationRole(req dto.CreateApplicationRoleRequest, createdBy uuid.UUID) (*dto.ApplicationRoleDTO, error) {
	db := config.DB

	appID, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		return nil, errors.New("application_id inválido")
	}

	// Verificar que la aplicación existe
	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	// Verificar si ya existe un rol con el mismo nombre en la misma aplicación
	var exists int64
	if err := db.Model(&models.ApplicationRole{}).
		Where("application_id = ? AND name = ? AND is_deleted = FALSE", appID, req.Name).
		Count(&exists).Error; err != nil {
		return nil, err
	}

	if exists > 0 {
		return nil, errors.New("ya existe un rol con este nombre en esta aplicación")
	}

	role := models.ApplicationRole{
		ID:            uuid.New(),
		Name:          req.Name,
		Description:   req.Description,
		ApplicationID: appID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsDeleted:     false,
	}

	if err := db.Create(&role).Error; err != nil {
		return nil, err
	}

	result := mapper.ToApplicationRoleDTO(role, &app, 0, 0)
	return &result, nil
}