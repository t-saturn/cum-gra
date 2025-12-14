package handlers

import (
	"server/internal/dto"
	services "server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetOrganicUnitsStatsHandler(c fiber.Ctx) error {
	result, err := services.GetOrganicUnitsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de unidades orgánicas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}