package services

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetRolesApp(pageStr, pageSizeStr, isDeletedStr string) (*dto.RolesAppResponse, error) {
	db := config.DB

	page := 1
	pageSize := 20
	isDeleted := false

	if v, err := strconv.Atoi(pageStr); err == nil && v > 0 {
		page = v
	}
	if v, err := strconv.Atoi(pageSizeStr); err == nil && v > 0 && v <= 200 {
		pageSize = v
	}
	if b, err := strconv.ParseBool(isDeletedStr); err == nil {
		isDeleted = b
	}

	var total int64
	totalQ := db.Model(&models.Application{}).Where("is_deleted = ?", isDeleted)
	if err := totalQ.Count(&total).Error; err != nil {
		return nil, err
	}

	var apps []models.Application
	if err := db.Model(&models.Application{}).
		Select("id", "name", "client_id").
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&apps).Error; err != nil {
		return nil, err
	}

	appDTOs := make([]dto.AppMinimalDTO, 0, len(apps))
	appIDs := make([]uuid.UUID, 0, len(apps))
	for _, a := range apps {
		appDTOs = append(appDTOs, dto.AppMinimalDTO{
			ID:       a.ID,
			Name:     a.Name,
			ClientID: a.ClientID,
		})
		appIDs = append(appIDs, a.ID)
	}

	if len(appIDs) == 0 {
		return &dto.RolesAppResponse{
			Data: dto.RolesAppData{
				Apps:    []dto.AppMinimalDTO{},
				Roles:   []dto.RoleMinimalDTO{},
				Modules: []dto.ModuleMinimalDTO{},
			},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	var roles []models.ApplicationRole
	if err := db.Model(&models.ApplicationRole{}).
		Select("id", "name").
		Where("application_id IN ?", appIDs).
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Find(&roles).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	roleDTOs := make([]dto.RoleMinimalDTO, 0, len(roles))
	for _, r := range roles {
		roleDTOs = append(roleDTOs, dto.RoleMinimalDTO{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	var modules []models.Module
	qModules := db.Model(&models.Module{}).
		Select("id", "name", "icon").
		Where("application_id IN ?", appIDs).
		Order("sort_order ASC, created_at DESC")

	if isDeleted {
		qModules = qModules.Where("deleted_at IS NOT NULL")
	} else {
		qModules = qModules.Where("deleted_at IS NULL")
	}

	if err := qModules.Find(&modules).Error; err != nil {
		return nil, err
	}

	moduleDTOs := make([]dto.ModuleMinimalDTO, 0, len(modules))
	for _, m := range modules {
		moduleDTOs = append(moduleDTOs, dto.ModuleMinimalDTO{
			ID:   m.ID,
			Name: m.Name,
			Icon: m.Icon,
		})
	}

	resp := &dto.RolesAppResponse{
		Data: dto.RolesAppData{
			Apps:    appDTOs,
			Roles:   roleDTOs,
			Modules: moduleDTOs,
		},
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	return resp, nil
}
