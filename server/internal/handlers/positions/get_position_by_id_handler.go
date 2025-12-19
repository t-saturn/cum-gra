package handlers

import (
	"server/internal/dto"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetStructuralPositionByIDHandler(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.GetStructuralPositionByID(id)
	if err != nil {
		if err.Error() == "posición estructural no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo posición estructural:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}