package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserModuleRestrictionByID(id string) (*dto.UserModuleRestrictionDTO, error) {
	db := config.DB

	restrictionID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var restriction models.UserModuleRestriction
	if err := db.Where("id = ?", restrictionID).First(&restriction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restricción no encontrada")
		}
		return nil, err
	}

	// Obtener usuario
	var user models.User
	if err := db.Where("id = ?", restriction.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	// Obtener user detail
	var userDetail models.UserDetail
	db.Where("user_id = ?", restriction.UserID).First(&userDetail)

	// Obtener módulo
	var module models.Module
	if err := db.Where("id = ?", restriction.ModuleID).First(&module).Error; err != nil {
		return nil, err
	}

	// Obtener application
	var app models.Application
	if err := db.Where("id = ?", restriction.ApplicationID).First(&app).Error; err != nil {
		return nil, err
	}

	result := mapper.ToUserModuleRestrictionDTO(restriction, &user, &userDetail, &module, &app)
	return &result, nil
}