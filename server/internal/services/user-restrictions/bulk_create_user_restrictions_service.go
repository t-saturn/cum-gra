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

func BulkCreateUserModuleRestrictions(req dto.BulkCreateUserModuleRestrictionsRequest, createdBy uuid.UUID) (*dto.BulkCreateUserModuleRestrictionsResponse, error) {
	db := config.DB

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("user_id inválido")
	}

	appID, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		return nil, errors.New("application_id inválido")
	}

	// Verificar que el usuario existe
	var user models.User
	if err := db.Where("id = ? AND is_deleted = FALSE", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Obtener user detail
	var userDetail models.UserDetail
	db.Where("user_id = ?", userID).First(&userDetail)

	// Verificar que la aplicación existe
	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	response := &dto.BulkCreateUserModuleRestrictionsResponse{
		Created: 0,
		Skipped: 0,
		Failed:  0,
		Details: []dto.UserModuleRestrictionDTO{},
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

		// Verificar que el módulo pertenece a la aplicación
		if module.ApplicationID == nil || *module.ApplicationID != appID {
			response.Failed++
			continue
		}

		// Verificar si ya existe
		var exists int64
		if err := db.Model(&models.UserModuleRestriction{}).
			Where("user_id = ? AND module_id = ? AND application_id = ? AND is_deleted = FALSE", userID, moduleID, appID).
			Count(&exists).Error; err != nil {
			response.Failed++
			continue
		}

		if exists > 0 {
			response.Skipped++
			continue
		}

		// Crear la restricción
		restriction := models.UserModuleRestriction{
			ID:              uuid.New(),
			UserID:          userID,
			ModuleID:        moduleID,
			ApplicationID:   appID,
			RestrictionType: req.RestrictionType,
			Reason:          req.Reason,
			CreatedAt:       time.Now(),
			CreatedBy:       createdBy,
			UpdatedAt:       time.Now(),
			IsDeleted:       false,
		}

		if err := db.Create(&restriction).Error; err != nil {
			response.Failed++
			continue
		}

		response.Created++
		response.Details = append(response.Details,
			mapper.ToUserModuleRestrictionDTO(restriction, &user, &userDetail, &module, &app))
	}

	return response, nil
}