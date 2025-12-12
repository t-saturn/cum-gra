package handlers

import (
	"server/internal/dto"
	"server/internal/services/applications"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func CreateApplicationHandler(c fiber.Ctx) error {
	var req dto.CreateApplicationRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

	// TODO: Validar request con validator
	// if err := validator.Validate(req); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	// Obtener user_id del contexto (del token de Keycloak)
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

	result, err := srvapplications.CreateApplication(req, userID)
	if err != nil {
		if err.Error() == "ya existe una aplicación con este client_id" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error creando aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
