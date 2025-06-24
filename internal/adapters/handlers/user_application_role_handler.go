package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/dto"
	validate "github.com/t-saturn/central-user-manager/pkg/validator"
)

type UserApplicationRoleHandler struct {
	service *services.UserApplicationRoleService
}

func NewUserApplicationRoleHandler(s *services.UserApplicationRoleService) *UserApplicationRoleHandler {
	return &UserApplicationRoleHandler{service: s}
}

func (h *UserApplicationRoleHandler) Create(c fiber.Ctx) error {
	var input dto.CreateUserApplicationRoleDTO

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Create(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Rol de usuario creado exitosamente"})
}

func (h *UserApplicationRoleHandler) GetAll(c fiber.Ctx) error {
	roles, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(roles)
}
