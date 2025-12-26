package handlers

import (
	"server/internal/dto"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetAllPositionsHandler(c fiber.Ctx) error {
	onlyActive := true
	if v := c.Query("only_active"); v == "false" {
		onlyActive = false
	}

	result, err := services.GetAllPositions(onlyActive)
	if err != nil {
		logger.Log.Error("Error obteniendo todas las posiciones:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}