package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func GetUserApplicationRoles(page, pageSize int, isDeleted bool, userID, applicationID, isRevoked *string) (*dto.UserApplicationRolesListResponse, error) {
	db := config.DB

	query := db.Model(&models.UserApplicationRole{}).Where("is_deleted = ?", isDeleted)

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

	if isRevoked != nil {
		if *isRevoked == "true" {
			query = query.Where("revoked_at IS NOT NULL")
		} else if *isRevoked == "false" {
			query = query.Where("revoked_at IS NULL")
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var assignments []models.UserApplicationRole
	if err := query.
		Order("granted_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&assignments).Error; err != nil {
		return nil, err
	}

	if len(assignments) == 0 {
		return &dto.UserApplicationRolesListResponse{
			Data:     []dto.UserApplicationRoleDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	// Obtener IDs únicos
	userIDs := make([]uuid.UUID, 0)
	appIDs := make([]uuid.UUID, 0)
	roleIDs := make([]uuid.UUID, 0)
	grantedByIDs := make([]uuid.UUID, 0)
	revokedByIDs := make([]uuid.UUID, 0)

	for _, a := range assignments {
		userIDs = append(userIDs, a.UserID)
		appIDs = append(appIDs, a.ApplicationID)
		roleIDs = append(roleIDs, a.ApplicationRoleID)
		grantedByIDs = append(grantedByIDs, a.GrantedBy)
		if a.RevokedBy != nil {
			revokedByIDs = append(revokedByIDs, *a.RevokedBy)
		}
	}

	// Obtener usuarios
	allUserIDs := append(userIDs, grantedByIDs...)
	allUserIDs = append(allUserIDs, revokedByIDs...)
	
	var users []models.User
	if err := db.Where("id IN ?", allUserIDs).Find(&users).Error; err != nil {
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

	// Obtener applications
	var apps []models.Application
	if err := db.Where("id IN ?", appIDs).Find(&apps).Error; err != nil {
		return nil, err
	}
	appMap := make(map[uuid.UUID]*models.Application)
	for i := range apps {
		appMap[apps[i].ID] = &apps[i]
	}

	// Obtener roles
	var roles []models.ApplicationRole
	if err := db.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return nil, err
	}
	roleMap := make(map[uuid.UUID]*models.ApplicationRole)
	for i := range roles {
		roleMap[roles[i].ID] = &roles[i]
	}

	out := make([]dto.UserApplicationRoleDTO, 0, len(assignments))
	for _, assignment := range assignments {
		user := userMap[assignment.UserID]
		userDetail := userDetailMap[assignment.UserID]
		app := appMap[assignment.ApplicationID]
		role := roleMap[assignment.ApplicationRoleID]
		grantedByUser := userMap[assignment.GrantedBy]
		
		var revokedByUser *models.User
		if assignment.RevokedBy != nil {
			revokedByUser = userMap[*assignment.RevokedBy]
		}

		out = append(out, mapper.ToUserApplicationRoleDTO(assignment, user, userDetail, app, role, grantedByUser, revokedByUser))
	}

	return &dto.UserApplicationRolesListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}