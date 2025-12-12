package handlers

import (
	"strconv"
	"strings"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetUsersRestrictionsHandler(c fiber.Ctx) error {
	db := config.DB

	page := 1
	if v := strings.TrimSpace(c.Query("page", "")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			page = n
		}
	}
	pageSize := 20
	if v := strings.TrimSpace(c.Query("page_size", "")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 200 {
			pageSize = n
		}
	}
	isDeletedStr := strings.ToLower(strings.TrimSpace(c.Query("is_deleted", "")))
	isDeleted := false
	if isDeletedStr == "true" || isDeletedStr == "1" || isDeletedStr == "t" {
		isDeleted = true
	}

	var total int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ?", isDeleted).
		Count(&total).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsHandler.countUsers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	var users []models.User
	if err := db.
		Model(&models.User{}).
		Where("is_deleted = ?", isDeleted).
		Select("id, first_name, last_name, email, dni").
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&users).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsHandler.findUsers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}
	if len(users) == 0 {
		return c.JSON(dto.RolesRestrictResponseDTO{
			Data:     []dto.RoleRestrictDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	}

	userIDs := make([]uuid.UUID, 0, len(users))
	for _, u := range users {
		userIDs = append(userIDs, u.ID)
	}

	rolesByUserApp := make(map[uuid.UUID]map[uuid.UUID]*dto.RoleMinimalDTO)
	appInfo := make(map[uuid.UUID]dto.AppMinimalDTO)

	type roleRow struct {
		UserID      uuid.UUID
		AppID       uuid.UUID
		AppName     string
		AppClientID string
		RoleID      uuid.UUID
		RoleName    string
	}

	var roleRows []roleRow
	if err := db.
		Model(&models.UserApplicationRole{}).
		Select(`
			user_application_roles.user_id        AS user_id,
			user_application_roles.application_id AS app_id,
			applications.name                     AS app_name,
			applications.client_id                AS app_client_id,
			application_roles.id                  AS role_id,
			application_roles.name                AS role_name
		`).
		Joins(`JOIN applications ON applications.id = user_application_roles.application_id AND applications.is_deleted = FALSE`).
		Joins(`JOIN application_roles ON application_roles.id = user_application_roles.application_role_id AND application_roles.is_deleted = FALSE`).
		Where("user_application_roles.is_deleted = FALSE").
		Where("user_application_roles.user_id IN ?", userIDs).
		Find(&roleRows).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsHandler.rolesQuery: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	for _, r := range roleRows {
		if _, ok := rolesByUserApp[r.UserID]; !ok {
			rolesByUserApp[r.UserID] = make(map[uuid.UUID]*dto.RoleMinimalDTO)
		}
		if _, exists := rolesByUserApp[r.UserID][r.AppID]; !exists {
			rolesByUserApp[r.UserID][r.AppID] = &dto.RoleMinimalDTO{
				ID:   r.RoleID,
				Name: r.RoleName,
			}
		}
		if _, ok := appInfo[r.AppID]; !ok {
			appInfo[r.AppID] = dto.AppMinimalDTO{
				ID:       r.AppID,
				Name:     r.AppName,
				ClientID: r.AppClientID,
			}
		}
	}

	restrictedModules := make(map[uuid.UUID]map[uuid.UUID][]dto.ModuleMinimalDTO)

	type restrictRow struct {
		UserID      uuid.UUID
		AppID       uuid.UUID
		AppName     string
		AppClientID string
		ModuleID    uuid.UUID
		ModuleName  string
		ModuleIcon  *string
	}

	var restrictRows []restrictRow
	if err := db.
		Model(&models.UserModuleRestriction{}).
		Select(`
			user_module_restrictions.user_id      AS user_id,
			user_module_restrictions.application_id AS app_id,
			applications.name                     AS app_name,
			applications.client_id                AS app_client_id,
			modules.id                            AS module_id,
			modules.name                          AS module_name,
			modules.icon                          AS module_icon
		`).
		Joins(`JOIN applications ON applications.id = user_module_restrictions.application_id AND applications.is_deleted = FALSE`).
		Joins(`JOIN modules ON modules.id = user_module_restrictions.module_id AND modules.deleted_at IS NULL`).
		Where("user_module_restrictions.is_deleted = FALSE").
		Where("user_module_restrictions.user_id IN ?", userIDs).
		Find(&restrictRows).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsHandler.restrictQuery: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	for _, rr := range restrictRows {
		if _, ok := appInfo[rr.AppID]; !ok {
			appInfo[rr.AppID] = dto.AppMinimalDTO{
				ID:       rr.AppID,
				Name:     rr.AppName,
				ClientID: rr.AppClientID,
			}
		}
		if _, ok := restrictedModules[rr.UserID]; !ok {
			restrictedModules[rr.UserID] = make(map[uuid.UUID][]dto.ModuleMinimalDTO)
		}
		restrictedModules[rr.UserID][rr.AppID] = append(restrictedModules[rr.UserID][rr.AppID], dto.ModuleMinimalDTO{
			ID:   rr.ModuleID,
			Name: rr.ModuleName,
			Icon: rr.ModuleIcon,
		})
	}

	appsPerUser := make(map[uuid.UUID]map[uuid.UUID]struct{})
	for userID, m := range rolesByUserApp {
		if _, ok := appsPerUser[userID]; !ok {
			appsPerUser[userID] = map[uuid.UUID]struct{}{}
		}
		for appID := range m {
			appsPerUser[userID][appID] = struct{}{}
		}
	}
	for userID, m := range restrictedModules {
		if _, ok := appsPerUser[userID]; !ok {
			appsPerUser[userID] = map[uuid.UUID]struct{}{}
		}
		for appID := range m {
			appsPerUser[userID][appID] = struct{}{}
		}
	}

	appIDSet := make(map[uuid.UUID]struct{})
	for _, appSet := range appsPerUser {
		for appID := range appSet {
			appIDSet[appID] = struct{}{}
		}
	}
	appIDs := make([]uuid.UUID, 0, len(appIDSet))
	for id := range appIDSet {
		appIDs = append(appIDs, id)
	}

	modulesByApp := make(map[uuid.UUID][]dto.ModuleMinimalDTO)
	if len(appIDs) > 0 {
		type modRow struct {
			AppID      uuid.UUID
			ModuleID   uuid.UUID
			ModuleName string
			ModuleIcon *string
		}
		var modRows []modRow
		if err := db.
			Model(&models.Module{}).
			Select(`
				modules.application_id AS app_id,
				modules.id             AS module_id,
				modules.name           AS module_name,
				modules.icon           AS module_icon
			`).
			Where("modules.application_id IN ?", appIDs).
			Where("modules.deleted_at IS NULL").
			Find(&modRows).Error; err != nil {
			logger.Log.Errorf("GetUserRestrictionsHandler.modulesByApp: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
		}
		for _, mr := range modRows {
			modulesByApp[mr.AppID] = append(modulesByApp[mr.AppID], dto.ModuleMinimalDTO{
				ID:   mr.ModuleID,
				Name: mr.ModuleName,
				Icon: mr.ModuleIcon,
			})
		}
	}

	data := make([]dto.RoleRestrictDTO, 0, len(users))

	for _, u := range users {
		userDTO := dto.UserMinimalDTO{
			ID:    u.ID,
			Email: u.Email,
			DNI:   u.DNI,
		}

		apps := []dto.UserAppAssignmentDTO{}

		appSet, hasApps := appsPerUser[u.ID]
		if hasApps {
			for appID := range appSet {
				appDTO, ok := appInfo[appID]
				if !ok {
					continue
				}

				var roleDTO *dto.RoleMinimalDTO
				if roles, ok := rolesByUserApp[u.ID]; ok {
					if r, ok2 := roles[appID]; ok2 {
						roleDTO = &dto.RoleMinimalDTO{ID: r.ID, Name: r.Name}
					}
				}

				appModules := modulesByApp[appID]

				var restrictMods []dto.ModuleMinimalDTO
				if rm, ok := restrictedModules[u.ID]; ok {
					if arr, ok2 := rm[appID]; ok2 {
						restrictMods = arr
					}
				}

				apps = append(apps, dto.UserAppAssignmentDTO{
					App:             appDTO,
					Role:            roleDTO,
					Modules:         appModules,
					ModulesRestrict: restrictMods,
				})
			}
		}

		data = append(data, dto.RoleRestrictDTO{
			User: userDTO,
			Apps: apps,
		})
	}

	resp := dto.RolesRestrictResponseDTO{
		Data:     data,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return c.JSON(resp)
}
