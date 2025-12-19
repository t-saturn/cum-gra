package handlers

import (
	"server/internal/dto"
	services "server/internal/services/user-application-roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func GetUserApplicationRoleByIDHandler(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.GetUserApplicationRoleByID(id)
	if err != nil {
		if err.Error() == "asignación de rol no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error obteniendo asignación de rol:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}