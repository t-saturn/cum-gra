package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/positions"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func DeleteStructuralPositionHandler(c fiber.Ctx) error {
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

	if err := services.DeleteStructuralPosition(id, user.ID); err != nil {
		if err.Error() == "posición estructural no encontrada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "no se puede eliminar una posición estructural que tiene usuarios asignados" {
			return c.Status(fiber.StatusBadRequest).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error eliminando posición estructural:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Posición estructural eliminada correctamente",
	})
}

func RestoreStructuralPositionHandler(c fiber.Ctx) error {
	id := c.Params("id")

	if err := services.RestoreStructuralPosition(id); err != nil {
		if err.Error() == "posición estructural no encontrada o no está eliminada" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error restaurando posición estructural:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Posición estructural restaurada correctamente",
	})
}