package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteUserModuleRestriction(id string, deletedBy uuid.UUID) error {
	db := config.DB

	restrictionID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var restriction models.UserModuleRestriction
	if err := db.Where("id = ? AND is_deleted = FALSE", restrictionID).First(&restriction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restricción no encontrada")
		}
		return err
	}

	now := time.Now()
	restriction.IsDeleted = true
	restriction.DeletedAt = &now
	restriction.DeletedBy = &deletedBy
	restriction.UpdatedAt = now

	return db.Save(&restriction).Error
}

func RestoreUserModuleRestriction(id string) error {
	db := config.DB

	restrictionID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var restriction models.UserModuleRestriction
	if err := db.Where("id = ? AND is_deleted = TRUE", restrictionID).First(&restriction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restricción no encontrada o no está eliminada")
		}
		return err
	}

	restriction.IsDeleted = false
	restriction.DeletedAt = nil
	restriction.DeletedBy = nil
	restriction.UpdatedAt = time.Now()

	return db.Save(&restriction).Error
}