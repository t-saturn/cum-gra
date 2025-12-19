package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
)

func CreateUbigeo(req dto.CreateUbigeoRequest) (*dto.UbigeoDTO, error) {
	db := config.DB

	// Verificar ubigeo_code único
	var exists int64
	if err := db.Model(&models.Ubigeo{}).
		Where("ubigeo_code = ?", req.UbigeoCode).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe un ubigeo con este código")
	}

	ubigeo := models.Ubigeo{
		UbigeoCode: req.UbigeoCode,
		IneiCode:   req.IneiCode,
		Department: req.Department,
		Province:   req.Province,
		District:   req.District,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := db.Create(&ubigeo).Error; err != nil {
		return nil, err
	}

	result := mapper.ToUbigeoDTO(ubigeo)
	return &result, nil
}