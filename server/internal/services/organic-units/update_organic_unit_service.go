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

func UpdateOrganicUnit(id string, req dto.UpdateOrganicUnitRequest, updatedBy uuid.UUID) (*dto.OrganicUnitItemDTO, error) {
	db := config.DB

	unitID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var unit models.OrganicUnit
	if err := db.Where("id = ? AND is_deleted = FALSE", uint(unitID)).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("unidad orgánica no encontrada")
		}
		return nil, err
	}

	// Validar nombre único si se está actualizando
	if req.Name != nil && *req.Name != unit.Name {
		var exists int64
		if err := db.Model(&models.OrganicUnit{}).
			Where("name = ? AND id != ? AND is_deleted = FALSE", *req.Name, uint(unitID)).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe una unidad orgánica con este nombre")
		}
		unit.Name = *req.Name
	}

	// Validar acronym único si se está actualizando
	if req.Acronym != nil && *req.Acronym != unit.Acronym {
		var exists int64
		if err := db.Model(&models.OrganicUnit{}).
			Where("acronym = ? AND id != ? AND is_deleted = FALSE", *req.Acronym, uint(unitID)).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe una unidad orgánica con este acrónimo")
		}
		unit.Acronym = *req.Acronym
	}

	// Validar parent_id si se está actualizando
	if req.ParentID != nil {
		if *req.ParentID == "" {
			unit.ParentID = nil
		} else {
			parentIDUint, err := strconv.ParseUint(*req.ParentID, 10, 32)
			if err != nil {
				return nil, errors.New("parent_id inválido")
			}

			// No puede ser su propio padre
			if uint(parentIDUint) == uint(unitID) {
				return nil, errors.New("una unidad orgánica no puede ser su propio padre")
			}

			var parent models.OrganicUnit
			if err := db.Where("id = ? AND is_deleted = FALSE", uint(parentIDUint)).First(&parent).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("unidad orgánica padre no encontrada")
				}
				return nil, err
			}
			pid := uint(parentIDUint)
			unit.ParentID = &pid
		}
	}

	if req.Brand != nil {
		unit.Brand = req.Brand
	}
	if req.Description != nil {
		unit.Description = req.Description
	}
	if req.IsActive != nil {
		unit.IsActive = *req.IsActive
	}
	if req.CodDepSGD != nil {
		unit.CodDepSGD = req.CodDepSGD
	}

	unit.UpdatedAt = time.Now()

	if err := db.Save(&unit).Error; err != nil {
		return nil, err
	}

	return GetOrganicUnitByID(id)
}