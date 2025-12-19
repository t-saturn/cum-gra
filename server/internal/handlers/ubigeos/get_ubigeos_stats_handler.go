package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUbigeosStatsHandler(c fiber.Ctx) error {
	result, err := services.GetUbigeosStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de ubigeos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}