package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateStructuralPosition(id string, req dto.UpdateStructuralPositionRequest, updatedBy uuid.UUID) (*dto.StructuralPositionItemDTO, error) {
	db := config.DB

	positionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var position models.StructuralPosition
	if err := db.Where("id = ? AND is_deleted = FALSE", uint(positionID)).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("posición estructural no encontrada")
		}
		return nil, err
	}

	// Validar nombre único si se está actualizando
	if req.Name != nil && *req.Name != position.Name {
		var exists int64
		if err := db.Model(&models.StructuralPosition{}).
			Where("name = ? AND id != ? AND is_deleted = FALSE", *req.Name, uint(positionID)).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe una posición estructural con este nombre")
		}
		position.Name = *req.Name
	}

	// Validar código único si se está actualizando
	if req.Code != nil && *req.Code != position.Code {
		var exists int64
		if err := db.Model(&models.StructuralPosition{}).
			Where("code = ? AND id != ? AND is_deleted = FALSE", *req.Code, uint(positionID)).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe una posición estructural con este código")
		}
		position.Code = *req.Code
	}

	if req.Level != nil {
		position.Level = req.Level
	}
	if req.Description != nil {
		position.Description = req.Description
	}
	if req.IsActive != nil {
		position.IsActive = *req.IsActive
	}
	if req.CodCarSGD != nil {
		position.CodCarSGD = req.CodCarSGD
	}

	position.UpdatedAt = time.Now()

	if err := db.Save(&position).Error; err != nil {
		return nil, err
	}

	return GetStructuralPositionByID(id)
}