package handlers

import (
	"server/internal/dto"
	"server/internal/services/applications"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func DeleteApplicationHandler(c fiber.Ctx) error {
	id := c.Params("id")

	// Obtener user_id del contexto
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no autenticado"})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "ID de usuario inválido"})
	}

	if err := services.DeleteApplication(id, userID); err != nil {
		if err.Error() == "aplicación no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Aplicación eliminada correctamente",
	})
}

func RestoreApplicationHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.RestoreApplication(id); err != nil {
		if err.Error() == "aplicación no encontrada o no está eliminada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error restaurando aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Aplicación restaurada correctamente",
	})
}
