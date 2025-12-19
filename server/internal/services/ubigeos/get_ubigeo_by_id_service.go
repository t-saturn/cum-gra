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

func GetUbigeoByID(id string) (*dto.UbigeoDTO, error) {
	db := config.DB

	ubigeoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var ubigeo models.Ubigeo
	if err := db.Where("id = ?", uint(ubigeoID)).First(&ubigeo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ubigeo no encontrado")
		}
		return nil, err
	}

	result := mapper.ToUbigeoDTO(ubigeo)
	return &result, nil
}