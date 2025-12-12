package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/modules"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func UpdateModuleHandler(c fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateModuleRequest
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

	result, err := services.UpdateModule(id, req, user.ID)
	if err != nil {
		if err.Error() == "módulo no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "ya existe un módulo con este nombre en esta aplicación" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "módulo padre no encontrado" ||
		   err.Error() == "parent_id inválido" ||
		   err.Error() == "un módulo no puede ser su propio padre" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error actualizando módulo:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}