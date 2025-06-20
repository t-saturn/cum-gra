package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"github.com/t-saturn/central-user-manager/pkg/validator"
	validate "github.com/t-saturn/central-user-manager/pkg/validator"
)

type StructuralPositionHandler struct {
	service *services.StructuralPositionService
}

func NewStructuralPositionHandler(s *services.StructuralPositionService) *StructuralPositionHandler {
	return &StructuralPositionHandler{service: s}
}

func (h *StructuralPositionHandler) Create(c fiber.Ctx) error {
	var input dto.CreateStructuralPositionDTO

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := validate.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Validación fallida",
			"detail": err.Error(),
		})
	}

	if err := h.service.Create(input); err != nil {
		logger.Log.Error("Error al crear el cargo:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Cargo registrado correctamente"})
}

func (h *StructuralPositionHandler) GetAll(c fiber.Ctx) error {
	structuralPositions, err := h.service.GetAll()
	if err != nil {
		logger.Log.Error("Error al obtener los cargos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": structuralPositions})
}

func (h *StructuralPositionHandler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	position, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cargo no encontrado"})
	}
	return c.JSON(position)
}

func (h *StructuralPositionHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var input dto.UpdateStructuralPositionDTO
	if e := c.Bind().Body(&input); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	if er := validator.Validate.Struct(input); er != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": er.Error()})
	}

	err = h.service.Update(id, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar"})
	}
	return c.JSON(fiber.Map{"message": "Actualizado correctamente"})
}

func (h *StructuralPositionHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	err = h.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar"})
	}
	return c.JSON(fiber.Map{"message": "Eliminado correctamente"})
}
