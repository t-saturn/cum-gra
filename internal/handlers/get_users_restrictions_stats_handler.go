package handlers

import (
	"time"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUsersRestrictionsStatsHandler(c fiber.Ctx) error {
	db := config.DB
	now := time.Now()

	var total int64
	if err := db.
		Model(&models.UserModuleRestriction{}).
		Count(&total).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsStatsHandler.total: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	// Activas: is_deleted = false AND (expires_at IS NULL OR expires_at > now)
	var active int64
	if err := db.
		Model(&models.UserModuleRestriction{}).
		Where("is_deleted = FALSE").
		Where(db.Where("expires_at IS NULL").Or("expires_at > ?", now)).
		Count(&active).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsStatsHandler.active: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	// Usuarios restringidos: DISTINCT user_id con restricciones activas
	var restrictedUsers int64
	if err := db.
		Model(&models.UserModuleRestriction{}).
		Where("is_deleted = FALSE").
		Where(db.Where("expires_at IS NULL").Or("expires_at > ?", now)).
		// Distinct count de user_id
		Distinct("user_id").
		Count(&restrictedUsers).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsStatsHandler.restrictedUsers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	// Eliminadas
	var deleted int64
	if err := db.
		Model(&models.UserModuleRestriction{}).
		Where("is_deleted = TRUE").
		Count(&deleted).Error; err != nil {
		logger.Log.Errorf("GetUserRestrictionsStatsHandler.deleted: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno del servidor"})
	}

	resp := dto.UserRestrictionsStatsDTO{
		TotalRestrictions:   total,
		ActiveRestrictions:  active,
		RestrictedUsers:     restrictedUsers,
		DeletedRestrictions: deleted,
	}
	return c.JSON(resp)
}
