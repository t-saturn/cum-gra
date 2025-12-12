package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/module-role-permissions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func DeleteModuleRolePermissionHandler(c fiber.Ctx) error {
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

	if err := services.DeleteModuleRolePermission(id, user.ID); err != nil {
		if err.Error() == "permiso no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando permiso:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Permiso eliminado correctamente",
	})
}

func RestoreModuleRolePermissionHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.RestoreModuleRolePermission(id); err != nil {
		if err.Error() == "permiso no encontrado o no está eliminado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error restaurando permiso:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Permiso restaurado correctamente",
	})
}