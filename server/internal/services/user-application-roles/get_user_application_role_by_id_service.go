package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserApplicationRoleByID(id string) (*dto.UserApplicationRoleDTO, error) {
	db := config.DB

	assignmentID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var assignment models.UserApplicationRole
	if err := db.Where("id = ?", assignmentID).First(&assignment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("asignación de rol no encontrada")
		}
		return nil, err
	}

	// Obtener usuario
	var user models.User
	if err := db.Where("id = ?", assignment.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	// Obtener user detail
	var userDetail models.UserDetail
	db.Where("user_id = ?", assignment.UserID).First(&userDetail)

	// Obtener application
	var app models.Application
	if err := db.Where("id = ?", assignment.ApplicationID).First(&app).Error; err != nil {
		return nil, err
	}

	// Obtener role
	var role models.ApplicationRole
	if err := db.Where("id = ?", assignment.ApplicationRoleID).First(&role).Error; err != nil {
		return nil, err
	}

	// Obtener granted by user
	var grantedByUser models.User
	if err := db.Where("id = ?", assignment.GrantedBy).First(&grantedByUser).Error; err != nil {
		return nil, err
	}

	// Obtener revoked by user si existe
	var revokedByUser *models.User
	if assignment.RevokedBy != nil {
		var rbu models.User
		if err := db.Where("id = ?", *assignment.RevokedBy).First(&rbu).Error; err == nil {
			revokedByUser = &rbu
		}
	}

	result := mapper.ToUserApplicationRoleDTO(assignment, &user, &userDetail, &app, &role, &grantedByUser, revokedByUser)
	return &result, nil
}