package handlers

import (
	"server/internal/dto"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetStructuralPositionsStatsHandler(c fiber.Ctx) error {
	result, err := services.GetStructuralPositionsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de posiciones estructurales:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}