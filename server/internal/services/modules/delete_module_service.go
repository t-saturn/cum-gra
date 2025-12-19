package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteModule(id string, deletedBy uuid.UUID) error {
	db := config.DB

	moduleID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var module models.Module
	if err := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("módulo no encontrado")
		}
		return err
	}

	// Verificar si tiene hijos
	var childrenCount int64
	if err := db.Model(&models.Module{}).
		Where("parent_id = ? AND deleted_at IS NULL", moduleID).
		Count(&childrenCount).Error; err != nil {
		return err
	}

	if childrenCount > 0 {
		return errors.New("no se puede eliminar un módulo que tiene submódulos")
	}

	now := time.Now()
	module.DeletedAt = &now
	module.DeletedBy = &deletedBy
	module.UpdatedAt = now

	return db.Save(&module).Error
}

func RestoreModule(id string) error {
	db := config.DB

	moduleID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var module models.Module
	if err := db.Where("id = ? AND deleted_at IS NOT NULL", moduleID).First(&module).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("módulo no encontrado o no está eliminado")
		}
		return err
	}

	module.DeletedAt = nil
	module.DeletedBy = nil
	module.UpdatedAt = time.Now()

	return db.Save(&module).Error
}