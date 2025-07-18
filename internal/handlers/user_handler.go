package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/dto"
	"github.com/t-saturn/central-user-manager/internal/services"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

// CreateUser maneja POST /users y crea un nuevo usuario.
func CreateUser(c fiber.Ctx) error {
	var input dto.CreateUserDTO

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}

	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	user, err := services.CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
