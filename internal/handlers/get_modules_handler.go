package handlers

import (
	"strconv"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetModulesHandler(c fiber.Ctx) error {
	db := config.DB
	// Parámetros (?page=1&page_size=20&is_deleted=true|false)

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

	var total int64
	totalQ := db.Model(&models.Module{})
	if isDeleted {
		totalQ = totalQ.Where("deleted_at IS NOT NULL")
	} else {
		totalQ = totalQ.Where("deleted_at IS NULL")
	}
	if err := totalQ.Count(&total).Error; err != nil {
		logger.Log.Error("Error contando módulos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var modules []models.Module
	dataQ := db.
		Preload("Application").
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize)

	if isDeleted {
		dataQ = dataQ.Where("deleted_at IS NOT NULL")
	} else {
		dataQ = dataQ.Where("deleted_at IS NULL")
	}

	if err := dataQ.Find(&modules).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusOK).JSON(dto.ModulesListResponse{
				Data:     []dto.ModuleWithAppDTO{},
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			})
		}
		logger.Log.Error("Error obteniendo módulos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	if len(modules) == 0 {
		return c.Status(fiber.StatusOK).JSON(dto.ModulesListResponse{
			Data:     []dto.ModuleWithAppDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
	}

	moduleIDs := make([]uuid.UUID, 0, len(modules))
	for _, m := range modules {
		moduleIDs = append(moduleIDs, m.ID)
	}

	type countRow struct {
		ModuleID   uuid.UUID `gorm:"column:module_id"`
		UsersCount int64     `gorm:"column:users_count"`
	}
	var counts []countRow

	if err := db.
		Table("module_role_permissions mrp").
		Select("mrp.module_id, COUNT(DISTINCT uar.user_id) AS users_count").
		Joins(`JOIN user_application_roles uar
					ON uar.application_role_id = mrp.application_role_id
					AND uar.is_deleted = FALSE
					AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u
					ON u.id = uar.user_id
					AND u.is_deleted = FALSE`).
		Where("mrp.is_deleted = FALSE AND mrp.module_id IN ?", moduleIDs).
		Group("mrp.module_id").
		Scan(&counts).Error; err != nil {

		logger.Log.Error("Error contando usuarios por módulo:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	usersCountByModule := make(map[uuid.UUID]int64, len(counts))
	for _, r := range counts {
		usersCountByModule[r.ModuleID] = r.UsersCount
	}

	out := make([]dto.ModuleWithAppDTO, 0, len(modules))
	for _, m := range modules {
		usersCount := usersCountByModule[m.ID]
		out = append(out, mapper.ToModuleWithAppDTO(m, usersCount))
	}

	return c.Status(fiber.StatusOK).JSON(dto.ModulesListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
