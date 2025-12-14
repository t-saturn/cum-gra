package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteUserApplicationRole(id string, deletedBy uuid.UUID) error {
	db := config.DB

	assignmentID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var assignment models.UserApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", assignmentID).First(&assignment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asignación de rol no encontrada")
		}
		return err
	}

	now := time.Now()
	assignment.IsDeleted = true
	assignment.DeletedAt = &now
	assignment.DeletedBy = &deletedBy
	assignment.UpdatedAt = now

	return db.Save(&assignment).Error
}

func UndeleteUserApplicationRole(id string) error {
	db := config.DB

	assignmentID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var assignment models.UserApplicationRole
	if err := db.Where("id = ? AND is_deleted = TRUE", assignmentID).First(&assignment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asignación de rol no encontrada o no está eliminada")
		}
		return err
	}

	assignment.IsDeleted = false
	assignment.DeletedAt = nil
	assignment.DeletedBy = nil
	assignment.UpdatedAt = time.Now()

	return db.Save(&assignment).Error
}