package handlers

import (
	"central-user-manager/internal/dto"
	"central-user-manager/internal/services"
	"central-user-manager/pkg/logger"
	"central-user-manager/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

func CreateUser(c fiber.Ctx) error {
	var input dto.CreateUserDTO

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "Datos mal formateados"})
	}

	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	if _, err := services.CreateUser(input); err != nil {
		logger.Log.Errorf("Error al crear usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "No se pudo crear el usuario"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.MessageResponse{Message: "Usuario creado exitosamente"})
}
