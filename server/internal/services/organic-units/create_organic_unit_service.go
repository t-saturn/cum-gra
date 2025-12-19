package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrganicUnit(req dto.CreateOrganicUnitRequest, createdBy uuid.UUID) (*dto.OrganicUnitItemDTO, error) {
	db := config.DB

	// Verificar nombre único
	var exists int64
	if err := db.Model(&models.OrganicUnit{}).
		Where("name = ? AND is_deleted = FALSE", req.Name).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe una unidad orgánica con este nombre")
	}

	// Verificar acronym único
	if err := db.Model(&models.OrganicUnit{}).
		Where("acronym = ? AND is_deleted = FALSE", req.Acronym).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe una unidad orgánica con este acrónimo")
	}

	// Validar parent_id si se proporciona
	var parentID *uint
	if req.ParentID != nil && *req.ParentID != "" {
		parentIDUint, err := strconv.ParseUint(*req.ParentID, 10, 32)
		if err != nil {
			return nil, errors.New("parent_id inválido")
		}

		var parent models.OrganicUnit
		if err := db.Where("id = ? AND is_deleted = FALSE", uint(parentIDUint)).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("unidad orgánica padre no encontrada")
			}
			return nil, err
		}
		pid := uint(parentIDUint)
		parentID = &pid
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	unit := models.OrganicUnit{
		Name:        req.Name,
		Acronym:     req.Acronym,
		Brand:       req.Brand,
		Description: req.Description,
		ParentID:    parentID,
		IsActive:    isActive,
		CodDepSGD:   req.CodDepSGD,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	if err := db.Create(&unit).Error; err != nil {
		return nil, err
	}

	result := mapper.ToOrganicUnitItemDTO(unit, 0)
	return &result, nil
}