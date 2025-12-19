package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

func UpdateUbigeo(id string, req dto.UpdateUbigeoRequest) (*dto.UbigeoDTO, error) {
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

	// Validar ubigeo_code único si se está actualizando
	if req.UbigeoCode != nil && *req.UbigeoCode != ubigeo.UbigeoCode {
		var exists int64
		if err := db.Model(&models.Ubigeo{}).
			Where("ubigeo_code = ? AND id != ?", *req.UbigeoCode, uint(ubigeoID)).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe un ubigeo con este código")
		}
		ubigeo.UbigeoCode = *req.UbigeoCode
	}

	if req.IneiCode != nil {
		ubigeo.IneiCode = *req.IneiCode
	}
	if req.Department != nil {
		ubigeo.Department = *req.Department
	}
	if req.Province != nil {
		ubigeo.Province = *req.Province
	}
	if req.District != nil {
		ubigeo.District = *req.District
	}

	ubigeo.UpdatedAt = time.Now()

	if err := db.Save(&ubigeo).Error; err != nil {
		return nil, err
	}

	return GetUbigeoByID(id)
}