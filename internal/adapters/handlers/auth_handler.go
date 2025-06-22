// internal/adapters/handlers/auth_handler.go
package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var input dto.LoginDTO
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid data"})
	}
	if err := validator.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID, err := h.service.Authenticate(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user_id": userID})
}
