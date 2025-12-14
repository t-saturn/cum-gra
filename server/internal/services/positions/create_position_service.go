package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func CreateStructuralPosition(req dto.CreateStructuralPositionRequest, createdBy uuid.UUID) (*dto.StructuralPositionItemDTO, error) {
	db := config.DB

	// Verificar nombre único
	var exists int64
	if err := db.Model(&models.StructuralPosition{}).
		Where("name = ? AND is_deleted = FALSE", req.Name).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe una posición estructural con este nombre")
	}

	// Verificar código único
	if err := db.Model(&models.StructuralPosition{}).
		Where("code = ? AND is_deleted = FALSE", req.Code).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe una posición estructural con este código")
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	position := models.StructuralPosition{
		Name:        req.Name,
		Code:        req.Code,
		Level:       req.Level,
		Description: req.Description,
		IsActive:    isActive,
		CodCarSGD:   req.CodCarSGD,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	if err := db.Create(&position).Error; err != nil {
		return nil, err
	}

	row := mapper.StructuralPositionToRow(position, 0)
	result := mapper.ToStructuralPositionItemDTO(row)
	return &result, nil
}