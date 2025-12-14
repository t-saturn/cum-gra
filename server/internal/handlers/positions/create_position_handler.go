package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func CreateStructuralPositionHandler(c fiber.Ctx) error {
	var req dto.CreateStructuralPositionRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

	// Obtener email del contexto
	email, ok := c.Locals("email").(string)
	if !ok || email == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no autenticado"})
	}

	// Buscar el usuario en la BD por email
	var user models.User
	db := config.DB
	if err := db.Where("email = ? AND is_deleted = FALSE", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no encontrado en el sistema"})
	}

	result, err := services.CreateStructuralPosition(req, user.ID)
	if err != nil {
		if err.Error() == "ya existe una posición estructural con este nombre" ||
		   err.Error() == "ya existe una posición estructural con este código" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error creando posición estructural:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}