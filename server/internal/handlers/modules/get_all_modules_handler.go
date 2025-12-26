package handlers

import (
	"server/internal/dto"
	services "server/internal/services/modules"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetAllModulesHandler(c fiber.Ctx) error {
	onlyActive := true
	if v := c.Query("only_active"); v == "false" {
		onlyActive = false
	}

	var applicationID *string
	if v := c.Query("application_id"); v != "" {
		applicationID = &v
	}

	result, err := services.GetAllModules(onlyActive, applicationID)
	if err != nil {
		logger.Log.Error("Error obteniendo todos los módulos:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}