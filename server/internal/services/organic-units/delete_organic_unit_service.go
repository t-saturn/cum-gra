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

func DeleteOrganicUnit(id string, deletedBy uuid.UUID) error {
	db := config.DB

	unitID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("ID inválido")
	}

	var unit models.OrganicUnit
	if err := db.Where("id = ? AND is_deleted = FALSE", uint(unitID)).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("unidad orgánica no encontrada")
		}
		return err
	}

	// Verificar si tiene hijos
	var childrenCount int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("parent_id = ? AND is_deleted = FALSE", uint(unitID)).
		Count(&childrenCount).Error; err != nil {
		return err
	}

	if childrenCount > 0 {
		return errors.New("no se puede eliminar una unidad orgánica que tiene sub-unidades")
	}

	// Verificar si tiene usuarios asignados
	var usersCount int64
	if err := db.Table("user_details").
		Where("organic_unit_id = ?", uint(unitID)).
		Count(&usersCount).Error; err != nil {
		return err
	}

	if usersCount > 0 {
		return errors.New("no se puede eliminar una unidad orgánica que tiene usuarios asignados")
	}

	now := time.Now()
	unit.IsDeleted = true
	unit.DeletedAt = &now
	unit.DeletedBy = &deletedBy
	unit.UpdatedAt = now

	return db.Save(&unit).Error
}

func RestoreOrganicUnit(id string) error {
	db := config.DB

	unitID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("ID inválido")
	}

	var unit models.OrganicUnit
	if err := db.Where("id = ? AND is_deleted = TRUE", uint(unitID)).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("unidad orgánica no encontrada o no está eliminada")
		}
		return err
	}

	unit.IsDeleted = false
	unit.DeletedAt = nil
	unit.DeletedBy = nil
	unit.UpdatedAt = time.Now()

	return db.Save(&unit).Error
}