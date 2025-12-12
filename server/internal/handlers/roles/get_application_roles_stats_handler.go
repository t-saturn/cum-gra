package handlers

import (
	"server/internal/dto"
	"server/internal/services/roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetApplicationRolesStatsHandler(c fiber.Ctx) error {
	result, err := services.GetApplicationRolesStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de roles:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}