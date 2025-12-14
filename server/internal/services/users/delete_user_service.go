package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteUser(id string, deletedBy uuid.UUID) error {
	db := config.DB

	userID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var user models.User
	if err := db.Where("id = ? AND is_deleted = FALSE", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuario no encontrado")
		}
		return err
	}

	now := time.Now()
	user.IsDeleted = true
	user.DeletedAt = &now
	user.DeletedBy = &deletedBy
	user.UpdatedAt = now

	return db.Save(&user).Error
}

func RestoreUser(id string) error {
	db := config.DB

	userID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	var user models.User
	if err := db.Where("id = ? AND is_deleted = TRUE", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuario no encontrado o no está eliminado")
		}
		return err
	}

	user.IsDeleted = false
	user.DeletedAt = nil
	user.DeletedBy = nil
	user.UpdatedAt = time.Now()

	return db.Save(&user).Error
}