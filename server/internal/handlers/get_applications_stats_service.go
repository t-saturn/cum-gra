package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetApplicationsStatsHandler(c fiber.Ctx) error {
	db := config.DB

	stats, err := services.GetApplicationsStats(db)
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de aplicaciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}
