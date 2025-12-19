package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
)

func GetModules(page, pageSize int, isDeleted bool, applicationID *string) (*dto.ModulesListResponse, error) {
	db := config.DB

	query := db.Model(&models.Module{}).Where("deleted_at IS NULL = ?", !isDeleted)

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

	var modules []models.Module
	if err := query.
		Order("sort_order ASC, created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&modules).Error; err != nil {
		return nil, err
	}

	if len(modules) == 0 {
		return &dto.ModulesListResponse{
			Data:     []dto.ModuleWithAppDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	moduleIDs := make([]uuid.UUID, 0, len(modules))
	appIDs := make(map[uuid.UUID]struct{})
	
	for _, m := range modules {
		moduleIDs = append(moduleIDs, m.ID)
		if m.ApplicationID != nil {
			appIDs[*m.ApplicationID] = struct{}{}
		}
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

	// Obtener count de usuarios por módulo (a través de module_role_permissions y user_application_roles)
	var userCounts []struct {
		ModuleID uuid.UUID `gorm:"column:module_id"`
		Count    int64     `gorm:"column:count"`
	}
	if err := db.Table("module_role_permissions mrp").
		Select("mrp.module_id, COUNT(DISTINCT uar.user_id) as count").
		Joins("JOIN user_application_roles uar ON uar.application_role_id = mrp.application_role_id AND uar.is_deleted = FALSE AND uar.revoked_at IS NULL").
		Where("mrp.is_deleted = FALSE AND mrp.module_id IN ?", moduleIDs).
		Group("mrp.module_id").
		Scan(&userCounts).Error; err != nil {
		return nil, err
	}

	userCountMap := make(map[uuid.UUID]int64)
	for _, uc := range userCounts {
		userCountMap[uc.ModuleID] = uc.Count
	}

	out := make([]dto.ModuleWithAppDTO, 0, len(modules))
	for _, module := range modules {
		var app *models.Application
		if module.ApplicationID != nil {
			app = appMap[*module.ApplicationID]
		}
		usersCount := userCountMap[module.ID]
		out = append(out, mapper.ToModuleWithAppDTO(module, app, usersCount))
	}

	return &dto.ModulesListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}