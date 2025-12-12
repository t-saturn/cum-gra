package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func GetModuleRolePermissions(page, pageSize int, isDeleted bool, moduleID, roleID *string) (*dto.ModuleRolePermissionsListResponse, error) {
	db := config.DB

	query := db.Model(&models.ModuleRolePermission{}).Where("is_deleted = ?", isDeleted)

	if moduleID != nil && *moduleID != "" {
		modUUID, err := uuid.Parse(*moduleID)
		if err != nil {
			return nil, err
		}
		query = query.Where("module_id = ?", modUUID)
	}

	if roleID != nil && *roleID != "" {
		roleUUID, err := uuid.Parse(*roleID)
		if err != nil {
			return nil, err
		}
		query = query.Where("application_role_id = ?", roleUUID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var permissions []models.ModuleRolePermission
	if err := query.
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&permissions).Error; err != nil {
		return nil, err
	}

	if len(permissions) == 0 {
		return &dto.ModuleRolePermissionsListResponse{
			Data:     []dto.ModuleRolePermissionDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	moduleIDs := make([]uuid.UUID, 0)
	roleIDs := make([]uuid.UUID, 0)
	appIDs := make(map[uuid.UUID]struct{})

	for _, p := range permissions {
		moduleIDs = append(moduleIDs, p.ModuleID)
		roleIDs = append(roleIDs, p.ApplicationRoleID)
	}

	// Obtener módulos
	var modules []models.Module
	if err := db.Where("id IN ?", moduleIDs).Find(&modules).Error; err != nil {
		return nil, err
	}
	moduleMap := make(map[uuid.UUID]*models.Module)
	for i := range modules {
		moduleMap[modules[i].ID] = &modules[i]
		if modules[i].ApplicationID != nil {
			appIDs[*modules[i].ApplicationID] = struct{}{}
		}
	}

	// Obtener roles
	var roles []models.ApplicationRole
	if err := db.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return nil, err
	}
	roleMap := make(map[uuid.UUID]*models.ApplicationRole)
	for i := range roles {
		roleMap[roles[i].ID] = &roles[i]
		appIDs[roles[i].ApplicationID] = struct{}{}
	}

	// Obtener applications
	appIDList := make([]uuid.UUID, 0, len(appIDs))
	for id := range appIDs {
		appIDList = append(appIDList, id)
	}

	var apps []models.Application
	if len(appIDList) > 0 {
		if err := db.Where("id IN ?", appIDList).Find(&apps).Error; err != nil {
			return nil, err
		}
	}
	appMap := make(map[uuid.UUID]*models.Application)
	for i := range apps {
		appMap[apps[i].ID] = &apps[i]
	}

	out := make([]dto.ModuleRolePermissionDTO, 0, len(permissions))
	for _, perm := range permissions {
		module := moduleMap[perm.ModuleID]
		role := roleMap[perm.ApplicationRoleID]
		
		var app *models.Application
		if role != nil {
			app = appMap[role.ApplicationID]
		}
		
		out = append(out, mapper.ToModuleRolePermissionDTO(perm, module, role, app))
	}

	return &dto.ModuleRolePermissionsListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}