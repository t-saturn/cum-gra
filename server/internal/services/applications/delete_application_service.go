package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteApplication(id string, deletedBy uuid.UUID) error {
	db := config.DB

	appID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("aplicación no encontrada")
		}
		return err
	}

	now := time.Now()
	app.IsDeleted = true
	app.DeletedAt = &now
	app.DeletedBy = &deletedBy
	app.UpdatedAt = now

	return db.Save(&app).Error
}

func RestoreApplication(id string) error {
	db := config.DB

	appID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var app models.Application
	if err := db.Where("id = ? AND is_deleted = TRUE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("aplicación no encontrada o no está eliminada")
		}
		return err
	}

	app.IsDeleted = false
	app.DeletedAt = nil
	app.DeletedBy = nil
	app.UpdatedAt = time.Now()

	return db.Save(&app).Error
}