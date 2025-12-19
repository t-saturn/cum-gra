package handlers

import (
	"server/internal/dto"
	services "server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetOrganicUnitByIDHandler(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.GetOrganicUnitByID(id)
	if err != nil {
		if err.Error() == "unidad orgánica no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo unidad orgánica:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}