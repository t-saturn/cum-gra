package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type ApplicationHandler struct {
	service *services.ApplicationService
}

func NewApplicationHandler(service *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		service: service,
	}
}

func (h *ApplicationHandler) Create() fiber.Handler {
	return func(c fiber.Ctx) error {
		var input dto.CreateApplicationDTO

		// pasear el cuerpo
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Error al parsear el cuerpo",
			})
		}

		// valdiar la estructura del DTO
		if err := validator.Validate.Struct(input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
		}

		// Verificar si el nombre ya está en uso
		if exists, err := h.service.IsnameTaken(input.Name); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al verificar duplicidad",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "El nombre ya se encuentra registrado",
			})
		}

		// Ejecutar creación
		if err := h.service.Create(c.Context(), &input); err != nil {
			logger.Log.Error("Error al crear la aplicación: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo crear la aplicación",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
			Message: "Aplicación creada exitosamente",
		})
	}
}
func (h *ApplicationHandler) GetByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		app, err := h.service.GetByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al obtener la aplicación",
			})
		}
		if app == nil {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "aplicación no encontrada",
			})
		}

		return c.Status(fiber.StatusOK).JSON(app)
	}
}

func (h *ApplicationHandler) Update() fiber.Handler {
	return func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		// Bind del cuerpo
		var input dto.UpdateApplicationDTO
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Cuerpo de la solicitud inválido",
			})
		}

		// Validación de campos
		if err := validator.Validate.Struct(input); err != nil {
			errors := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: errors,
			})
		}

		// Validar si el nombre ya existe (si se envió)
		if input.Name != nil {
			exists, err := h.service.IsNameTakenExceptID(*input.Name, id)
			if err != nil {
				logger.Log.Error("Error al verificar nombre de aplicación: ", err)
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error interno al verificar nombre",
				})
			}
			if exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "Ya existe otra aplicación con este nombre",
				})
			}
		}

		// Ejecutar la actualización
		if err := h.service.Update(c.Context(), id, &input); err != nil {
			logger.Log.Error("Error al actualizar aplicación: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(dto.MessageResponse{
			Message: "Aplicación actualizada exitosamente",
		})
	}
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type ApplicationRoleHanlder struct {
	service *services.ApplicationRoleService
}

func NewApplicationRoleHanlder(service *services.ApplicationRoleService) *ApplicationRoleHanlder {
	return &ApplicationRoleHanlder{
		service: service,
	}
}

func (h *ApplicationRoleHanlder) Create() fiber.Handler {
	return func(c fiber.Ctx) error {
		var input dto.CreateApplicationRoleDTO

		// pasear el cuerpo
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Error al parsear el cuerpo",
			})
		}

		// valdiar la estructura del DTO
		if err := validator.Validate.Struct(input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
		}

		// Verificar si el nombre ya está en uso
		if exists, err := h.service.IsnameTaken(input.Name); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al verificar duplicidad",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "El nombre ya se encuentra registrado",
			})
		}

		// Ejecutar creación
		if err := h.service.Create(c.Context(), &input); err != nil {
			logger.Log.Error("Error al crear la aplicación: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo crear la aplicación",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
			Message: "Aplicación creada exitosamente",
		})
	}
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
