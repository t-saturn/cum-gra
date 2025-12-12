package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteApplicationRole(id string, deletedBy uuid.UUID) error {
	db := config.DB

	roleID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("rol no encontrado")
		}
		return err
	}

	now := time.Now()
	role.IsDeleted = true
	role.DeletedAt = &now
	role.DeletedBy = &deletedBy
	role.UpdatedAt = now

	return db.Save(&role).Error
}

func RestoreApplicationRole(id string) error {
	db := config.DB

	roleID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = TRUE", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("rol no encontrado o no está eliminado")
		}
		return err
	}

	role.IsDeleted = false
	role.DeletedAt = nil
	role.DeletedBy = nil
	role.UpdatedAt = time.Now()

	return db.Save(&role).Error
}