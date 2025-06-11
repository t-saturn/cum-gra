package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/dto"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
	"github.com/t-saturn/central-user-manager/server/internal/core/services"
)

var userService = services.NewUserService(repositories.NewUserRepository())

func CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "datos inválidos"})
	}

	user := &entities.User{
		Name:     req.Name,
		LastName: req.LastName,
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := userService.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "no se pudo registrar el usuario"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "usuario creado correctamente",
	})
}
