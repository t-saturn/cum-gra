package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/module-role-permissions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func CreateModuleRolePermissionHandler(c fiber.Ctx) error {
	var req dto.CreateModuleRolePermissionRequest
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

	result, err := services.CreateModuleRolePermission(req, user.ID)
	if err != nil {
		if err.Error() == "ya existe un permiso para este módulo y rol" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "módulo no encontrado" ||
		   err.Error() == "rol no encontrado" ||
		   err.Error() == "module_id inválido" ||
		   err.Error() == "application_role_id inválido" ||
		   err.Error() == "el módulo y el rol deben pertenecer a la misma aplicación" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error creando permiso:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}