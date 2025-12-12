package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteModuleRolePermission(id string, deletedBy uuid.UUID) error {
	db := config.DB

	permID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var permission models.ModuleRolePermission
	if err := db.Where("id = ? AND is_deleted = FALSE", permID).First(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("permiso no encontrado")
		}
		return err
	}

	now := time.Now()
	permission.IsDeleted = true
	permission.DeletedAt = &now
	permission.DeletedBy = &deletedBy

	return db.Save(&permission).Error
}

func RestoreModuleRolePermission(id string) error {
	db := config.DB

	permID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var permission models.ModuleRolePermission
	if err := db.Where("id = ? AND is_deleted = TRUE", permID).First(&permission).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("permiso no encontrado o no está eliminado")
		}
		return err
	}

	permission.IsDeleted = false
	permission.DeletedAt = nil
	permission.DeletedBy = nil

	return db.Save(&permission).Error
}