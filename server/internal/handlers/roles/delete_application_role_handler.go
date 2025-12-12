package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/services/roles"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func DeleteApplicationRoleHandler(c fiber.Ctx) error {
	id := c.Params("id")

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

	if err := services.DeleteApplicationRole(id, user.ID); err != nil {
		if err.Error() == "rol no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando rol:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Rol eliminado correctamente",
	})
}

func RestoreApplicationRoleHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.RestoreApplicationRole(id); err != nil {
		if err.Error() == "rol no encontrado o no está eliminado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error restaurando rol:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Rol restaurado correctamente",
	})
}