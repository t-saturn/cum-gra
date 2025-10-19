package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetOrganicUnitsStatsHandler(c fiber.Ctx) error {
	db := config.DB

	stats, err := services.GetOrganicUnitsStats(db)
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de unidades orgánicas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}
