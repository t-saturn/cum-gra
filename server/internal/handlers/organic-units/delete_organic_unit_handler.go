package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/organic-units"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func DeleteOrganicUnitHandler(c fiber.Ctx) error {
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

	if err := services.DeleteOrganicUnit(id, user.ID); err != nil {
		if err.Error() == "unidad orgánica no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "no se puede eliminar una unidad orgánica que tiene sub-unidades" ||
		   err.Error() == "no se puede eliminar una unidad orgánica que tiene usuarios asignados" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando unidad orgánica:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Unidad orgánica eliminada correctamente",
	})
}

func RestoreOrganicUnitHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.RestoreOrganicUnit(id); err != nil {
		if err.Error() == "unidad orgánica no encontrada o no está eliminada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error restaurando unidad orgánica:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Unidad orgánica restaurada correctamente",
	})
}