package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
)

type GetRolesAppOptions struct {
	Page      int
	PageSize  int
	IsDeleted *bool
}

func GetRolesAppMatrix(opts GetRolesAppOptions) (*dto.RolesAppResponseDTO, error) {
	db := config.DB

	page := opts.Page
	if page <= 0 {
		page = 1
	}
	pageSize := opts.PageSize
	if pageSize <= 0 || pageSize > 200 {
		pageSize = 20
	}

	countQ := db.Model(&models.ApplicationRole{})
	if opts.IsDeleted != nil {
		countQ = countQ.Where("is_deleted = ?", *opts.IsDeleted)
	}
	var total int64
	if err := countQ.Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		return &dto.RolesAppResponseDTO{
			Data:     []dto.RoleAppModulesItemDTO{},
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	type roleRow struct {
		RoleID      uuid.UUID
		RoleName    string
		AppID       uuid.UUID
		AppName     string
		AppClientID string
	}

	roleRowsQ := db.
		Model(&models.ApplicationRole{}).
		Select(`
			application_roles.id   AS role_id,
			application_roles.name AS role_name,
			applications.id        AS app_id,
			applications.name      AS app_name,
			applications.client_id AS app_client_id
		`).
		Joins(`JOIN applications ON applications.id = application_roles.application_id`).
		Order("application_roles.created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if opts.IsDeleted != nil {
		roleRowsQ = roleRowsQ.Where("application_roles.is_deleted = ?", *opts.IsDeleted)
	}

	var roleRows []roleRow
	if err := roleRowsQ.Find(&roleRows).Error; err != nil {
		return nil, err
	}
	if len(roleRows) == 0 {
		return &dto.RolesAppResponseDTO{
			Data:     []dto.RoleAppModulesItemDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	appIDSet := map[uuid.UUID]struct{}{}
	roleIDSet := map[uuid.UUID]struct{}{}
	for _, r := range roleRows {
		appIDSet[r.AppID] = struct{}{}
		roleIDSet[r.RoleID] = struct{}{}
	}
	appIDs := make([]uuid.UUID, 0, len(appIDSet))
	for id := range appIDSet {
		appIDs = append(appIDs, id)
	}
	roleIDs := make([]uuid.UUID, 0, len(roleIDSet))
	for id := range roleIDSet {
		roleIDs = append(roleIDs, id)
	}

	type appModRow struct {
		AppID      uuid.UUID
		ModuleID   uuid.UUID
		ModuleName string
		ModuleIcon *string
	}
	var appModRows []appModRow
	if err := db.
		Model(&models.Module{}).
		Select(`
			modules.application_id AS app_id,
			modules.id             AS module_id,
			modules.name           AS module_name,
			modules.icon           AS module_icon
		`).
		Where("modules.deleted_at IS NULL").
		Where("modules.application_id IN ?", appIDs).
		Find(&appModRows).Error; err != nil {
		return nil, err
	}

	appModules := make(map[uuid.UUID][]dto.ModuleMinimalDTO)
	for _, am := range appModRows {
		appModules[am.AppID] = append(appModules[am.AppID], dto.ModuleMinimalDTO{
			ID:   am.ModuleID,
			Name: am.ModuleName,
			Icon: am.ModuleIcon,
		})
	}

	type roleModRow struct {
		RoleID     uuid.UUID
		ModuleID   uuid.UUID
		ModuleName string
		ModuleIcon *string
	}
	var roleModRows []roleModRow
	if len(roleIDs) > 0 {
		if err := db.
			Table("module_role_permissions mrp").
			Select(`
				mrp.application_role_id AS role_id,
				m.id                     AS module_id,
				m.name                   AS module_name,
				m.icon                   AS module_icon
			`).
			Joins(`JOIN modules m ON m.id = mrp.module_id AND m.deleted_at IS NULL`).
			Joins(`JOIN application_roles ar ON ar.id = mrp.application_role_id AND ar.is_deleted = FALSE`).
			Where("mrp.is_deleted = FALSE").
			Where("mrp.application_role_id IN ?", roleIDs).
			Find(&roleModRows).Error; err != nil {
			return nil, err
		}
	}

	roleModules := make(map[uuid.UUID][]dto.ModuleMinimalDTO)
	for _, rm := range roleModRows {
		roleModules[rm.RoleID] = append(roleModules[rm.RoleID], dto.ModuleMinimalDTO{
			ID:   rm.ModuleID,
			Name: rm.ModuleName,
			Icon: rm.ModuleIcon,
		})
	}

	items := make([]dto.RoleAppModulesItemDTO, 0, len(roleRows))
	for _, r := range roleRows {
		items = append(items, dto.RoleAppModulesItemDTO{
			Role: dto.RoleMinimalDTO{
				ID:   r.RoleID,
				Name: r.RoleName,
			},
			App: dto.AppMinimalDTO{
				ID:       r.AppID,
				Name:     r.AppName,
				ClientID: r.AppClientID,
			},
			AppModules:  appModules[r.AppID],
			RoleModules: roleModules[r.RoleID],
		})
	}

	return &dto.RolesAppResponseDTO{
		Data:     items,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
