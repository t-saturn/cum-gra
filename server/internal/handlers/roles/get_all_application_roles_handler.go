package handlers

import (
	"server/internal/dto"
	services "server/internal/services/roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetAllApplicationRolesHandler(c fiber.Ctx) error {
	var applicationID *string
	if v := c.Query("application_id"); v != "" {
		applicationID = &v
	}

	result, err := services.GetAllApplicationRoles(applicationID)
	if err != nil {
		logger.Log.Error("Error obteniendo todos los roles:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}