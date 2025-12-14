package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/user-restrictions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func CreateUserModuleRestrictionHandler(c fiber.Ctx) error {
	var req dto.CreateUserModuleRestrictionRequest
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

	result, err := services.CreateUserModuleRestriction(req, user.ID)
	if err != nil {
		if err.Error() == "ya existe una restricción activa para este usuario y módulo" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "usuario no encontrado" ||
		   err.Error() == "módulo no encontrado" ||
		   err.Error() == "aplicación no encontrada" ||
		   err.Error() == "el módulo no pertenece a la aplicación especificada" ||
		   err.Error() == "user_id inválido" ||
		   err.Error() == "module_id inválido" ||
		   err.Error() == "application_id inválido" ||
		   err.Error() == "formato de fecha inválido para expires_at" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error creando restricción:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}