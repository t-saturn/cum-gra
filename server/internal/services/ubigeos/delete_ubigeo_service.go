package services

import (
	"errors"
	"strconv"

	"server/internal/config"
	"server/internal/models"

	"gorm.io/gorm"
)

func DeleteUbigeo(id string) error {
	db := config.DB

	ubigeoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("ID inválido")
	}

	var ubigeo models.Ubigeo
	if err := db.Where("id = ?", uint(ubigeoID)).First(&ubigeo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("ubigeo no encontrado")
		}
		return err
	}

	// Verificar si hay usuarios asignados a este ubigeo
	var usersCount int64
	if err := db.Table("user_details").
		Where("ubigeo_id = ?", uint(ubigeoID)).
		Count(&usersCount).Error; err != nil {
		return err
	}

	if usersCount > 0 {
		return errors.New("no se puede eliminar un ubigeo que tiene usuarios asignados")
	}

	return db.Delete(&ubigeo).Error
}