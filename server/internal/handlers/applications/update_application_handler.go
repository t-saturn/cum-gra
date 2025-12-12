package handlers

import (
	"server/internal/dto"
	"server/internal/services/applications"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func UpdateApplicationHandler(c fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateApplicationRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

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

	result, err := srvapplications.UpdateApplication(id, req, userID)
	if err != nil {
		if err.Error() == "aplicación no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "ya existe una aplicación con este client_id" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error actualizando aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
