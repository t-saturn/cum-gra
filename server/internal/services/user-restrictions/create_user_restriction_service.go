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

func CreateUserModuleRestriction(req dto.CreateUserModuleRestrictionRequest, createdBy uuid.UUID) (*dto.UserModuleRestrictionDTO, error) {
	db := config.DB

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("user_id inválido")
	}

	moduleID, err := uuid.Parse(req.ModuleID)
	if err != nil {
		return nil, errors.New("module_id inválido")
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

	// Verificar que el módulo existe
	var module models.Module
	if err := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("módulo no encontrado")
		}
		return nil, err
	}

	// Verificar que la aplicación existe
	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	// Verificar que el módulo pertenece a la aplicación
	if module.ApplicationID == nil || *module.ApplicationID != appID {
		return nil, errors.New("el módulo no pertenece a la aplicación especificada")
	}

	// Verificar si ya existe una restricción activa
	var exists int64
	if err := db.Model(&models.UserModuleRestriction{}).
		Where("user_id = ? AND module_id = ? AND application_id = ? AND is_deleted = FALSE", userID, moduleID, appID).
		Count(&exists).Error; err != nil {
		return nil, err
	}

	if exists > 0 {
		return nil, errors.New("ya existe una restricción activa para este usuario y módulo")
	}

	var expiresAt *time.Time
	if req.ExpiresAt != nil {
		parsed, err := time.Parse(time.RFC3339, *req.ExpiresAt)
		if err != nil {
			return nil, errors.New("formato de fecha inválido para expires_at")
		}
		expiresAt = &parsed
	}

	restriction := models.UserModuleRestriction{
		ID:                 uuid.New(),
		UserID:             userID,
		ModuleID:           moduleID,
		ApplicationID:      appID,
		RestrictionType:    req.RestrictionType,
		MaxPermissionLevel: req.MaxPermissionLevel,
		Reason:             req.Reason,
		ExpiresAt:          expiresAt,
		CreatedAt:          time.Now(),
		CreatedBy:          createdBy,
		UpdatedAt:          time.Now(),
		IsDeleted:          false,
	}

	if err := db.Create(&restriction).Error; err != nil {
		return nil, err
	}

	result := mapper.ToUserModuleRestrictionDTO(restriction, &user, &userDetail, &module, &app)
	return &result, nil
}