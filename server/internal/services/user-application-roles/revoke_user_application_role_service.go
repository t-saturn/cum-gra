package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RevokeUserApplicationRole(id string, revokedBy uuid.UUID) error {
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

	if assignment.RevokedAt != nil {
		return errors.New("este rol ya fue revocado anteriormente")
	}

	now := time.Now()
	assignment.RevokedAt = &now
	assignment.RevokedBy = &revokedBy
	assignment.UpdatedAt = now

	return db.Save(&assignment).Error
}

func RestoreUserApplicationRole(id string) error {
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

	if assignment.RevokedAt == nil {
		return errors.New("este rol no ha sido revocado")
	}

	assignment.RevokedAt = nil
	assignment.RevokedBy = nil
	assignment.UpdatedAt = time.Now()

	return db.Save(&assignment).Error
}