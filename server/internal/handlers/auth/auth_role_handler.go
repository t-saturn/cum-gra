package handlers

import (
	"server/internal/dto"
	"server/internal/services/auth"
	"server/pkg/logger"
	"server/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AuthRoleHandler(c fiber.Ctx) error {
	var input dto.AuthRoleRequest

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "Datos mal formateados"})
	}
	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "UserID inválido"})
	}

	resp, err := services.GetUserRoleAndModules(userID, input.ClientID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "No se encontraron datos"})
		}
		logger.Log.Error("Error en AuthService:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
