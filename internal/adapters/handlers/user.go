package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
		if err := h.service.Create(c.Context(), &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo crear la posición estructural",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
			Message: "Posición estructural creada exitosamente",
		})
	}
}

func (h *StructuralPositionHandler) GetByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Obtener el ID de la ruta
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		// Buscar por ID
		result, err := h.service.GetByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al buscar la posición estructural",
			})
		}

		if result == nil {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "Posición estructural no encontrada",
			})
		}

		return c.JSON(result)
	}
}

func (h *StructuralPositionHandler) Update() fiber.Handler {
	return func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		// Obtener datos desde query params
		input := dto.UpdateStructuralPositionDTO{}

		if name := c.Query("name"); name != "" {
			input.Name = &name
		}

		if code := c.Query("code"); code != "" {
			input.Code = &code
		}

		if level := c.Query("level"); level != "" {
			levelInt := 0
			if levelParsed, err := strconv.Atoi(level); err == nil {
				levelInt = levelParsed
			}
			input.Level = &levelInt
		}

		if desc := c.Query("description"); desc != "" {
			input.Description = &desc
		}

		if active := c.Query("is_active"); active != "" {
			val, err := strconv.ParseBool(active)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "El parámetro 'is_active' debe ser true o false",
				})
			}
			input.IsActive = &val
		}

		// Validación personalizada solo si se intenta actualizar valores únicos
		if input.Name != nil {
			exists, err := h.service.IsNameTakenExceptID(*input.Name, id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar nombre",
				})
			}
			if exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "Ya existe otra posición estructural con este nombre",
				})
			}
		}

		if input.Code != nil {
			exists, err := h.service.IsCodeTakenExceptID(*input.Code, id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar código",
				})
			}
			if exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "Ya existe otra posición estructural con este código",
				})
			}
		}

		if err := h.service.Update(c.Context(), id, &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo actualizar la posición estructural",
			})
		}

		return c.Status(fiber.StatusOK).JSON(dto.MessageResponse{
			Message: "Posición estructural actualizada exitosamente",
		})
	}
}
