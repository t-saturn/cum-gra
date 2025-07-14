package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

type StructuralPositionHandler struct {
	service *services.StructuralPositionService
}

func NewStructuralPositionHandler(service *services.StructuralPositionService) *StructuralPositionHandler {
	return &StructuralPositionHandler{
		service: service,
	}
}

func (h *StructuralPositionHandler) Create() fiber.Handler {
	return func(c fiber.Ctx) error {
		var input dto.CreateStructuralPositionDTO

		// Parsear el body
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "cuerpo de la solicitud inválido",
			})
		}

		// Validar el DTO
		if err := validator.Validate.Struct(input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
		}

		if exists, err := h.service.IsNameTaken(input.Name); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Error al verificar nombre",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Ya existe una posición estructural con este nombre",
			})
		}

		if exists, err := h.service.IsCodeTaken(input.Code); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Error al verificar código",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Ya existe una posición estructural con este código",
			})
		}

		// Llamar a servicio y retornar
		_, err := h.service.Create(c.Context(), &input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
			Message: "Posición estructural creada exitosamente",
		})
	}
}
