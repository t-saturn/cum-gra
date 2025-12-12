package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

type RoleCountRow struct {
	RoleID       uuid.UUID `gorm:"column:role_id"`
	ModulesCount int64     `gorm:"column:modules_count"`
	UsersCount   int64     `gorm:"column:users_count"`
}

func GetApplicationRoles(page, pageSize int, isDeleted bool, applicationID *string) (*dto.ApplicationRolesListResponse, error) {
	db := config.DB

	query := db.Model(&models.ApplicationRole{}).Where("is_deleted = ?", isDeleted)

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

	var roles []models.ApplicationRole
	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&roles).Error; err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return &dto.ApplicationRolesListResponse{
			Data:     []dto.ApplicationRoleDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	roleIDs := make([]uuid.UUID, 0, len(roles))
	appIDs := make(map[uuid.UUID]struct{})
	for _, r := range roles {
		roleIDs = append(roleIDs, r.ID)
		appIDs[r.ApplicationID] = struct{}{}
	}

	// Obtener applications
	appIDList := make([]uuid.UUID, 0, len(appIDs))
	for id := range appIDs {
		appIDList = append(appIDList, id)
	}

	var apps []models.Application
	if err := db.Where("id IN ?", appIDList).Find(&apps).Error; err != nil {
		return nil, err
	}

	appMap := make(map[uuid.UUID]*models.Application)
	for i := range apps {
		appMap[apps[i].ID] = &apps[i]
	}

	// Obtener count de módulos por rol
	var moduleCounts []struct {
		RoleID uuid.UUID `gorm:"column:application_role_id"`
		Count  int64     `gorm:"column:count"`
	}
	if err := db.Table("module_role_permissions").
		Select("application_role_id, COUNT(*) as count").
		Where("is_deleted = FALSE AND application_role_id IN ?", roleIDs).
		Group("application_role_id").
		Scan(&moduleCounts).Error; err != nil {
		return nil, err
	}

	moduleCountMap := make(map[uuid.UUID]int64)
	for _, mc := range moduleCounts {
		moduleCountMap[mc.RoleID] = mc.Count
	}

	// Obtener count de usuarios por rol
	var userCounts []struct {
		RoleID uuid.UUID `gorm:"column:application_role_id"`
		Count  int64     `gorm:"column:count"`
	}
	if err := db.Table("user_application_roles").
		Select("application_role_id, COUNT(DISTINCT user_id) as count").
		Where("is_deleted = FALSE AND revoked_at IS NULL AND application_role_id IN ?", roleIDs).
		Group("application_role_id").
		Scan(&userCounts).Error; err != nil {
		return nil, err
	}

	userCountMap := make(map[uuid.UUID]int64)
	for _, uc := range userCounts {
		userCountMap[uc.RoleID] = uc.Count
	}

	out := make([]dto.ApplicationRoleDTO, 0, len(roles))
	for _, role := range roles {
		app := appMap[role.ApplicationID]
		modulesCount := moduleCountMap[role.ID]
		usersCount := userCountMap[role.ID]
		out = append(out, mapper.ToApplicationRoleDTO(role, app, modulesCount, usersCount))
	}

	return &dto.ApplicationRolesListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}