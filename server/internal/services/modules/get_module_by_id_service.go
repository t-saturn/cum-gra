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

func GetModuleByID(id string) (*dto.ModuleWithAppDTO, error) {
	db := config.DB

	moduleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var module models.Module
	if err := db.Where("id = ?", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("módulo no encontrado")
		}
		return nil, err
	}

	// Obtener application si existe
	var app *models.Application
	if module.ApplicationID != nil {
		var appModel models.Application
		if err := db.Where("id = ?", *module.ApplicationID).First(&appModel).Error; err == nil {
			app = &appModel
		}
	}

	// Obtener count de usuarios
	var usersCount int64
	if err := db.Table("module_role_permissions mrp").
		Joins("JOIN user_application_roles uar ON uar.application_role_id = mrp.application_role_id AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL").
		Where("mrp.is_deleted = FALSE AND mrp.module_id = ?", moduleID).
		Distinct("uar.user_id").
		Count(&usersCount).Error; err != nil {
		return nil, err
	}

	result := mapper.ToModuleWithAppDTO(module, app, usersCount)
	return &result, nil
}