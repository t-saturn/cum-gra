package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateUserModuleRestriction(id string, req dto.UpdateUserModuleRestrictionRequest, updatedBy uuid.UUID) (*dto.UserModuleRestrictionDTO, error) {
	db := config.DB

	restrictionID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var restriction models.UserModuleRestriction
	if err := db.Where("id = ? AND is_deleted = FALSE", restrictionID).First(&restriction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restricción no encontrada")
		}
		return nil, err
	}

	if req.RestrictionType != nil {
		restriction.RestrictionType = *req.RestrictionType
	}

	if req.MaxPermissionLevel != nil {
		restriction.MaxPermissionLevel = req.MaxPermissionLevel
	}

	if req.Reason != nil {
		restriction.Reason = req.Reason
	}

	if req.ExpiresAt != nil {
		if *req.ExpiresAt == "" {
			restriction.ExpiresAt = nil
		} else {
			parsed, err := time.Parse(time.RFC3339, *req.ExpiresAt)
			if err != nil {
				return nil, errors.New("formato de fecha inválido para expires_at")
			}
			restriction.ExpiresAt = &parsed
		}
	}

	restriction.UpdatedAt = time.Now()
	restriction.UpdatedBy = &updatedBy

	if err := db.Save(&restriction).Error; err != nil {
		return nil, err
	}

	return GetUserModuleRestrictionByID(id)
}