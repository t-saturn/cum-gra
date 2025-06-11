package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/dto"
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
)

func CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "datos inválidos"})
	}

	user := entities.User{
		ID:        uuid.New(),
		Name:      req.Name,
		LastName:  req.LastName,
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  req.Password, // NOTA: En producción se debe hashear
	}

	// Aquí iría la llamada a service (omitiendo por simplicidad)

	return c.Status(fiber.StatusCreated).JSON(dto.UserResponse{
		ID:       user.ID.String(),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	})
}
