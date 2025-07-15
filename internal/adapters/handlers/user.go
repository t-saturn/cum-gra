package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
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
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "El parámetro 'level' debe ser un número entero",
				})
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

		// Validar el DTO si hay datos
		if err := validator.Validate.Struct(&input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
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

		// Ejecutar actualización
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

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type OrganicUnitHandler struct {
	service *services.OrganicUnitService
}

func NewOrganicUnitHandler(service *services.OrganicUnitService) *OrganicUnitHandler {
	return &OrganicUnitHandler{
		service: service,
	}
}

func (h *OrganicUnitHandler) Create() fiber.Handler {
	return func(c fiber.Ctx) error {
		var input dto.CreateOrganicUnitDTO

		// Parsear el cuerpo
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Cuerpo de la solicitud inválido",
			})
		}

		// Validar estructura
		if err := validator.Validate.Struct(input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
		}

		// Verificar nombre duplicado
		if exists, err := h.service.IsNameTaken(input.Name); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al verificar nombre",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Ya existe una unidad orgánica con este nombre",
			})
		}

		// Verificar sigla duplicada
		if exists, err := h.service.IsAcronymTaken(input.Acronym); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al verificar sigla",
			})
		} else if exists {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Ya existe una unidad orgánica con esta sigla",
			})
		}

		// Verificar si el parent_id existe (si se envió)
		if input.ParentID != nil {
			exists, err := h.service.IsIdTaken(*input.ParentID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar unidad padre",
				})
			}
			if !exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "La unidad orgánica padre especificada no existe",
				})
			}
		}

		// Crear
		if err := h.service.Create(c.Context(), &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo crear la unidad orgánica",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
			Message: "Unidad orgánica creada exitosamente",
		})
	}
}

func (h *OrganicUnitHandler) GetByID() fiber.Handler {
	return func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		organicUnit, err := h.service.GetByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error al obtener la unidad orgánica",
			})
		}
		if organicUnit == nil {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "Unidad orgánica no encontrada",
			})
		}

		return c.Status(fiber.StatusOK).JSON(organicUnit)
	}
}

func (h *OrganicUnitHandler) Update() fiber.Handler {
	return func(c fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "ID inválido",
			})
		}

		input := dto.UpdateOrganicUnitDTO{}

		if name := c.Query("name"); name != "" {
			input.Name = &name
		}
		if acronym := c.Query("acronym"); acronym != "" {
			input.Acronym = &acronym
		}
		if brand := c.Query("brand"); brand != "" {
			input.Brand = &brand
		}
		if desc := c.Query("description"); desc != "" {
			input.Description = &desc
		}

		// Validar el DTO si se enviaron campos
		if err := validator.Validate.Struct(&input); err != nil {
			translated := validator.FormatValidationError(err)
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Errors: translated,
			})
		}

		if parentID := c.Query("parent_id"); parentID != "" {
			parsedID, err := uuid.Parse(parentID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "parent_id inválido",
				})
			}
			if parsedID == id {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "La unidad orgánica no puede ser su propio padre",
				})
			}
			exists, err := h.service.IsIdTakenExeptedID(parsedID.String(), id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar el ID del padre",
				})
			}
			if !exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "La unidad orgánica padre no existe",
				})
			}
			input.ParentID = &parsedID
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

		if input.Name != nil {
			exists, err := h.service.IsNameTakenExceptID(*input.Name, id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar nombre",
				})
			}
			if exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "Ya existe otra unidad orgánica con este nombre",
				})
			}
		}

		if input.Acronym != nil {
			exists, err := h.service.IsAcronymTakenExceptID(*input.Acronym, id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
					Error: "Error al verificar acrónimo",
				})
			}
			if exists {
				return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
					Error: "Ya existe otra unidad orgánica con este acrónimo",
				})
			}
		}

		if err := h.service.Update(c.Context(), id, &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "No se pudo actualizar la unidad orgánica",
			})
		}

		return c.Status(fiber.StatusOK).JSON(dto.MessageResponse{
			Message: "Unidad orgánica actualizada exitosamente",
		})
	}
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
