// Package handlers define los controladores HTTP que gestionan las peticiones del cliente.
package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/dto"
	"github.com/t-saturn/central-user-manager/internal/services"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

// CreateUser maneja la solicitud POST /users, valida la entrada y crea un nuevo usuario en el sistema.
func CreateUser(c fiber.Ctx) error {
	var input dto.CreateUserDTO

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Datos mal formateados",
		})
	}

	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	if _, err := services.CreateUser(input); err != nil {
		logger.Log.Errorf("Error al crear usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "No se pudo crear el usuario",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{
		Message: "Usuario creado exitosamente",
	})
}
