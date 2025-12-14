package services

import (
	"errors"
	"strconv"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"gorm.io/gorm"
)

func GetOrganicUnitByID(id string) (*dto.OrganicUnitItemDTO, error) {
	db := config.DB

	unitID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var unit models.OrganicUnit
	if err := db.Where("id = ?", uint(unitID)).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("unidad orgánica no encontrada")
		}
		return nil, err
	}

	// Obtener count de usuarios
	var usersCount int64
	if err := db.Table("user_details").
		Where("organic_unit_id = ?", uint(unitID)).
		Count(&usersCount).Error; err != nil {
		return nil, err
	}

	result := mapper.ToOrganicUnitItemDTO(unit, usersCount)
	return &result, nil
}