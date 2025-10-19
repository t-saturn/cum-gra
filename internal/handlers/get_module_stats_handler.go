package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetModulesStatsHandler(c fiber.Ctx) error {
	db := config.DB

	var totalModules int64
	if err := db.Model(&models.Module{}).
		Count(&totalModules).Error; err != nil {
		logger.Log.Error("Error contando total de módulos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var activeModules int64
	if err := db.Model(&models.Module{}).
		Where("status = ? AND deleted_at IS NULL", "active").
		Count(&activeModules).Error; err != nil {
		logger.Log.Error("Error contando módulos activos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var deletedModules int64
	if err := db.Model(&models.Module{}).
		Where("deleted_at IS NOT NULL").
		Count(&deletedModules).Error; err != nil {
		logger.Log.Error("Error contando módulos eliminados:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var totalUsers int64
	if err := db.
		Table("module_role_permissions mrp").
		Joins(`JOIN modules m ON m.id = mrp.module_id AND m.deleted_at IS NULL`).
		Joins(`JOIN user_application_roles uar
					ON uar.application_role_id = mrp.application_role_id
					AND uar.is_deleted = FALSE
					AND uar.revoked_at IS NULL`).
		Joins(`JOIN users u
					ON u.id = uar.user_id
					AND u.is_deleted = FALSE`).
		Where("mrp.is_deleted = FALSE").
		Select("COUNT(DISTINCT u.id)").
		Scan(&totalUsers).Error; err != nil {

		logger.Log.Error("Error contando usuarios únicos con acceso a módulos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ModulesStatsResponse{
		TotalModules:   totalModules,
		ActiveModules:  activeModules,
		DeletedModules: deletedModules,
		TotalUsers:     totalUsers,
	})
}
