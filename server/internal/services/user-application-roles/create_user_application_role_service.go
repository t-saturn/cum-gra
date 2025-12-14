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

func CreateUserApplicationRole(req dto.CreateUserApplicationRoleRequest, grantedBy uuid.UUID) (*dto.UserApplicationRoleDTO, error) {
	db := config.DB

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("user_id inválido")
	}

	appID, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		return nil, errors.New("application_id inválido")
	}

	roleID, err := uuid.Parse(req.ApplicationRoleID)
	if err != nil {
		return nil, errors.New("application_role_id inválido")
	}

	// Verificar que el usuario existe
	var user models.User
	if err := db.Where("id = ? AND is_deleted = FALSE", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Obtener user detail
	var userDetail models.UserDetail
	db.Where("user_id = ?", userID).First(&userDetail)

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

	// Verificar si ya existe una asignación activa (no revocada y no eliminada)
	var exists int64
	if err := db.Model(&models.UserApplicationRole{}).
		Where("user_id = ? AND application_id = ? AND application_role_id = ? AND is_deleted = FALSE AND revoked_at IS NULL", 
			userID, appID, roleID).
		Count(&exists).Error; err != nil {
		return nil, err
	}

	if exists > 0 {
		return nil, errors.New("el usuario ya tiene este rol asignado en esta aplicación")
	}

	// Obtener granted by user
	var grantedByUser models.User
	if err := db.Where("id = ?", grantedBy).First(&grantedByUser).Error; err != nil {
		return nil, err
	}

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
		return nil, err
	}

	result := mapper.ToUserApplicationRoleDTO(assignment, &user, &userDetail, &app, &role, &grantedByUser, nil)
	return &result, nil
}