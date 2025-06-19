package handlers

import (
	"github.com/central-user-manager/internal/core/services"
	"github.com/central-user-manager/internal/shared/dto"
	"github.com/central-user-manager/pkg/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type OrganicUnitHandler struct {
	service *services.OrganicUnitService
}

func NewOrganicUnitHandler(s *services.OrganicUnitService) *OrganicUnitHandler {
	return &OrganicUnitHandler{service: s}
}

func (h *OrganicUnitHandler) Create(c fiber.Ctx) error {
	var input dto.CreateOrganicUnitDTO
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.Create(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Unidad registrada"})
}

func (h *OrganicUnitHandler) GetAll(c fiber.Ctx) error {
	units, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al listar"})
	}
	return c.JSON(units)
}

func (h *OrganicUnitHandler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	unit, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unidad no encontrada"})
	}
	return c.JSON(unit)
}

func (h *OrganicUnitHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var input dto.UpdateOrganicUnitDTO
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Update(id, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar"})
	}

	return c.JSON(fiber.Map{"message": "Actualizado correctamente"})
}

func (h *OrganicUnitHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar"})
	}
	return c.JSON(fiber.Map{"message": "Eliminado correctamente"})
}
