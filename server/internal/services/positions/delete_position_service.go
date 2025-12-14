package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteStructuralPosition(id string, deletedBy uuid.UUID) error {
	db := config.DB

	positionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("ID inválido")
	}

	var position models.StructuralPosition
	if err := db.Where("id = ? AND is_deleted = FALSE", uint(positionID)).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("posición estructural no encontrada")
		}
		return err
	}

	// Verificar si tiene usuarios asignados
	var usersCount int64
	if err := db.Table("user_details").
		Where("structural_position_id = ?", uint(positionID)).
		Count(&usersCount).Error; err != nil {
		return err
	}

	if usersCount > 0 {
		return errors.New("no se puede eliminar una posición estructural que tiene usuarios asignados")
	}

	now := time.Now()
	position.IsDeleted = true
	position.DeletedAt = &now
	position.DeletedBy = &deletedBy
	position.UpdatedAt = now

	return db.Save(&position).Error
}

func RestoreStructuralPosition(id string) error {
	db := config.DB

	positionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("ID inválido")
	}

	var position models.StructuralPosition
	if err := db.Where("id = ? AND is_deleted = TRUE", uint(positionID)).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("posición estructural no encontrada o no está eliminada")
		}
		return err
	}

	position.IsDeleted = false
	position.DeletedAt = nil
	position.DeletedBy = nil
	position.UpdatedAt = time.Now()

	return db.Save(&position).Error
}