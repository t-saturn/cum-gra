package handlers

import (
	"server/internal/dto"
	services "server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetAllOrganicUnitsHandler(c fiber.Ctx) error {
	onlyActive := true
	if v := c.Query("only_active"); v == "false" {
		onlyActive = false
	}

	result, err := services.GetAllOrganicUnits(onlyActive)
	if err != nil {
		logger.Log.Error("Error obteniendo todas las unidades orgánicas:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}