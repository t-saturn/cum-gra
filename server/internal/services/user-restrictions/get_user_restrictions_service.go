package services

import (
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func GetUserModuleRestrictions(page, pageSize int, isDeleted bool, userID, applicationID *string) (*dto.UserModuleRestrictionsListResponse, error) {
	db := config.DB

	query := db.Model(&models.UserModuleRestriction{}).Where("is_deleted = ?", isDeleted)

	if userID != nil && *userID != "" {
		userUUID, err := uuid.Parse(*userID)
		if err != nil {
			return nil, err
		}
		query = query.Where("user_id = ?", userUUID)
	}

	if applicationID != nil && *applicationID != "" {
		appUUID, err := uuid.Parse(*applicationID)
		if err != nil {
			return nil, err
		}
		query = query.Where("application_id = ?", appUUID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var restrictions []models.UserModuleRestriction
	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&restrictions).Error; err != nil {
		return nil, err
	}

	if len(restrictions) == 0 {
		return &dto.UserModuleRestrictionsListResponse{
			Data:     []dto.UserModuleRestrictionDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	// Obtener IDs únicos
	userIDs := make([]uuid.UUID, 0)
	moduleIDs := make([]uuid.UUID, 0)
	appIDs := make([]uuid.UUID, 0)

	for _, r := range restrictions {
		userIDs = append(userIDs, r.UserID)
		moduleIDs = append(moduleIDs, r.ModuleID)
		appIDs = append(appIDs, r.ApplicationID)
	}

	// Obtener usuarios
	var users []models.User
	if err := db.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}
	userMap := make(map[uuid.UUID]*models.User)
	for i := range users {
		userMap[users[i].ID] = &users[i]
	}

	// Obtener user details
	var userDetails []models.UserDetail
	if err := db.Where("user_id IN ?", userIDs).Find(&userDetails).Error; err != nil {
		return nil, err
	}
	userDetailMap := make(map[uuid.UUID]*models.UserDetail)
	for i := range userDetails {
		userDetailMap[userDetails[i].UserID] = &userDetails[i]
	}

	// Obtener módulos
	var modules []models.Module
	if err := db.Where("id IN ?", moduleIDs).Find(&modules).Error; err != nil {
		return nil, err
	}
	moduleMap := make(map[uuid.UUID]*models.Module)
	for i := range modules {
		moduleMap[modules[i].ID] = &modules[i]
	}

	// Obtener applications
	var apps []models.Application
	if err := db.Where("id IN ?", appIDs).Find(&apps).Error; err != nil {
		return nil, err
	}
	appMap := make(map[uuid.UUID]*models.Application)
	for i := range apps {
		appMap[apps[i].ID] = &apps[i]
	}

	out := make([]dto.UserModuleRestrictionDTO, 0, len(restrictions))
	now := time.Now()
	for _, restriction := range restrictions {
		// Verificar si la restricción ha expirado
		if restriction.ExpiresAt != nil && restriction.ExpiresAt.Before(now) {
			continue // Saltar restricciones expiradas
		}

		user := userMap[restriction.UserID]
		userDetail := userDetailMap[restriction.UserID]
		module := moduleMap[restriction.ModuleID]
		app := appMap[restriction.ApplicationID]

		out = append(out, mapper.ToUserModuleRestrictionDTO(restriction, user, userDetail, module, app))
	}

	return &dto.UserModuleRestrictionsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}