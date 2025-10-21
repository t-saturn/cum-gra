package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetRolesAppHandler(c fiber.Ctx) error {
	db := config.DB

	page := 1
	pageSize := 20
	if v := c.Query("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			page = n
		}
	}
	if v := c.Query("page_size"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 200 {
			pageSize = n
		}
	}
	isDeleted := false
	if v := c.Query("is_deleted"); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			isDeleted = b
		}
	}

	var apps []models.Application
	qApps := db.Model(&models.Application{}).
		Select("id", "name", "client_id").
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := qApps.Find(&apps).Error; err != nil {
		logger.Log.Error("GetRolesAppHandler.apps", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch applications",
		})
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
		resp := dto.RolesAppResponse{
			Apps:    []dto.AppMinimalDTO{},
			Roles:   []dto.RoleMinimalDTO{},
			Modules: []dto.ModuleMinimalDTO{},
		}
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	var roles []models.ApplicationRole
	qRoles := db.Model(&models.ApplicationRole{}).
		Select("id", "name").
		Where("application_id IN ?", appIDs).
		Where("is_deleted = ?", isDeleted).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if err := qRoles.Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			roles = []models.ApplicationRole{}
		} else {
			logger.Log.Error("GetRolesAppHandler.roles", "err", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to fetch roles",
			})
		}
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
		Order("sort_order ASC, created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if isDeleted {
		qModules = qModules.Where("deleted_at IS NOT NULL")
	} else {
		qModules = qModules.Where("deleted_at IS NULL")
	}

	if err := qModules.Find(&modules).Error; err != nil {
		logger.Log.Error("GetRolesAppHandler.modules", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch modules",
		})
	}

	moduleDTOs := make([]dto.ModuleMinimalDTO, 0, len(modules))
	for _, m := range modules {
		moduleDTOs = append(moduleDTOs, dto.ModuleMinimalDTO{
			ID:   m.ID,
			Name: m.Name,
			Icon: m.Icon,
		})
	}

	resp := dto.RolesAppResponse{
		Apps:    appDTOs,
		Roles:   roleDTOs,
		Modules: moduleDTOs,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
