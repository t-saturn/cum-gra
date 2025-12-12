package handlers

import (
	"server/internal/dto"
	services "server/internal/services/modules"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetModulesStatsHandler(c fiber.Ctx) error {
	result, err := services.GetModulesStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de módulos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}