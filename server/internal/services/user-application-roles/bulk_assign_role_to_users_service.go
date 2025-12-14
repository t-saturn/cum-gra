package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func BulkAssignRoleToUsers(req dto.BulkAssignRoleToUsersRequest, grantedBy uuid.UUID) (*dto.BulkAssignRoleToUsersResponse, error) {
	db := config.DB

	appID, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		return nil, errors.New("application_id inválido")
	}

	roleID, err := uuid.Parse(req.ApplicationRoleID)
	if err != nil {
		return nil, errors.New("application_role_id inválido")
	}

	// Verificar que la aplicación existe
	var app models.Application
	if err := db.Where("id = ? AND is_deleted = FALSE", appID).First(&app).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("aplicación no encontrada")
		}
		return nil, err
	}

	// Verificar que el rol existe
	var role models.ApplicationRole
	if err := db.Where("id = ? AND is_deleted = FALSE", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rol no encontrado")
		}
		return nil, err
	}

	// Verificar que el rol pertenece a la aplicación
	if role.ApplicationID != appID {
		return nil, errors.New("el rol no pertenece a la aplicación especificada")
	}

	// Obtener granted by user
	var grantedByUser models.User
	if err := db.Where("id = ?", grantedBy).First(&grantedByUser).Error; err != nil {
		return nil, err
	}

	response := &dto.BulkAssignRoleToUsersResponse{
		Created: 0,
		Skipped: 0,
		Failed:  0,
		Details: []dto.UserApplicationRoleDTO{},
	}

	for _, userIDStr := range req.UserIDs {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			response.Failed++
			continue
		}

		// Verificar que el usuario existe
		var user models.User
		if err := db.Where("id = ? AND is_deleted = FALSE", userID).First(&user).Error; err != nil {
			response.Failed++
			continue
		}

		// Obtener user detail
		var userDetail models.UserDetail
		db.Where("user_id = ?", userID).First(&userDetail)

		// Verificar si ya existe
		var exists int64
		if err := db.Model(&models.UserApplicationRole{}).
			Where("user_id = ? AND application_id = ? AND application_role_id = ? AND is_deleted = FALSE AND revoked_at IS NULL",
				userID, appID, roleID).
			Count(&exists).Error; err != nil {
			response.Failed++
			continue
		}

		if exists > 0 {
			response.Skipped++
			continue
		}

		// Crear la asignación
		assignment := models.UserApplicationRole{
			ID:                uuid.New(),
			UserID:            userID,
			ApplicationID:     appID,
			ApplicationRoleID: roleID,
			GrantedAt:         time.Now(),
			GrantedBy:         grantedBy,
			IsDeleted:         false,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		if err := db.Create(&assignment).Error; err != nil {
			response.Failed++
			continue
		}

		response.Created++
		response.Details = append(response.Details,
			mapper.ToUserApplicationRoleDTO(assignment, &user, &userDetail, &app, &role, &grantedByUser, nil))
	}

	return response, nil
}