package handlers

import (
	"server/internal/dto"
	services "server/internal/services/user-restrictions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUserModuleRestrictionsStatsHandler(c fiber.Ctx) error {
	result, err := services.GetUserModuleRestrictionsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de restricciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}