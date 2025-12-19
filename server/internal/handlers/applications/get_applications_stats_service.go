package handlers

import (
	"server/internal/dto"
	"server/internal/services/applications"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetApplicationsStatsHandler(c fiber.Ctx) error {

	stats, err := services.GetApplicationsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de aplicaciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}
