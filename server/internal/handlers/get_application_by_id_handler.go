package handlers

import (
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetApplicationByIDHandler(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.GetApplicationByID(id)
	if err != nil {
		if err.Error() == "aplicación no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}