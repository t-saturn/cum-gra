package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/dto"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres/repositories"
	"github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"
	"github.com/t-saturn/central-user-manager/server/internal/core/services"
	"github.com/t-saturn/central-user-manager/server/pkg/utils"
)

var userService = services.NewUserService(repositories.NewUserRepository())

func CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "datos inválidos"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "no se pudo encriptar la contraseña"})
	}

	user := &entities.User{
		Name:     req.Name,
		LastName: req.LastName,
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := userService.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "no se pudo registrar el usuario"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "usuario creado correctamente",
	})
}
