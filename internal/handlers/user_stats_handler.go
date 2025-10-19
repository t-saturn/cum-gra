package handlers

import (
	"time"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUsersStatsHandler(c fiber.Ctx) error {
	db := config.DB

	// Punto de corte para "último mes" (rolling)
	since := time.Now().AddDate(0, -1, 0)

	var total int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ?", false).
		Count(&total).Error; err != nil {
		logger.Log.Error("Error contando total de usuarios:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var active int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ? AND status = ?", false, "active").
		Count(&active).Error; err != nil {
		logger.Log.Error("Error contando usuarios activos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var suspended int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ? AND status = ?", false, "inactive").
		Count(&suspended).Error; err != nil {
		logger.Log.Error("Error contando usuarios suspendidos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var newLastMonth int64
	if err := db.Model(&models.User{}).
		Where("is_deleted = ? AND created_at >= ?", false, since).
		Count(&newLastMonth).Error; err != nil {
		logger.Log.Error("Error contando usuarios nuevos del último mes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.UsersStatsResponse{
		TotalUsers:        total,
		ActiveUsers:       active,
		SuspendedUsers:    suspended,
		NewUsersLastMonth: newLastMonth,
	})
}
