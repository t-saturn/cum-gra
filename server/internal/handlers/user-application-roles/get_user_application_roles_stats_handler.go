package handlers

import (
	"server/internal/dto"
	services "server/internal/services/user-application-roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUserApplicationRolesStatsHandler(c fiber.Ctx) error {
	result, err := services.GetUserApplicationRolesStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de asignaciones de roles:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}