package handlers

import (
	"github.com/central-user-manager/internal/core/services"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/central-user-manager/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type StructuralPositionHandler struct {
	service *services.StructuralPositionService
}

func NewStructuralPositionHandler(s *services.StructuralPositionService) *StructuralPositionHandler {
	return &StructuralPositionHandler{service: s}
}

func (h *StructuralPositionHandler) Create(c *fiber.Ctx) error {
	var input dto.CreateStructuralPositionDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := h.service.Create(input); err != nil {
		logger.Log.Error("Error al crear el cargo:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Cargo registrado correctamente"})
}
