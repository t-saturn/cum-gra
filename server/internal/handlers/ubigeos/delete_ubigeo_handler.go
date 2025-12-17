package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func DeleteUbigeoHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.DeleteUbigeo(id); err != nil {
		if err.Error() == "ubigeo no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "no se puede eliminar un ubigeo que tiene usuarios asignados" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando ubigeo:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ubigeo eliminado correctamente",
	})
}