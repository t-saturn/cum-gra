package handlers

import (
	"server/internal/dto"
	services "server/internal/services/users"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUserByIDHandler(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.GetUserByID(id)
	if err != nil {
		if err.Error() == "usuario no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo usuario:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}