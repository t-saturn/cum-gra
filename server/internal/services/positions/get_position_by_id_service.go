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

func GetStructuralPositionByID(id string) (*dto.StructuralPositionItemDTO, error) {
	db := config.DB

	positionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var position models.StructuralPosition
	if err := db.Where("id = ?", uint(positionID)).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("posición estructural no encontrada")
		}
		return nil, err
	}

	// Obtener count de usuarios
	var usersCount int64
	if err := db.Table("user_details").
		Where("structural_position_id = ?", uint(positionID)).
		Count(&usersCount).Error; err != nil {
		return nil, err
	}

	row := mapper.StructuralPositionToRow(position, usersCount)
	result := mapper.ToStructuralPositionItemDTO(row)
	return &result, nil
}