package handlers

import (
	"server/internal/dto"
	services "server/internal/services/module-role-permissions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetModuleRolePermissionsStatsHandler(c fiber.Ctx) error {
	result, err := services.GetModuleRolePermissionsStats()
	if err != nil {
		logger.Log.Error("Error obteniendo estadísticas de permisos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}