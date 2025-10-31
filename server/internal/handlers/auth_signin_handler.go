package handlers

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/pkg/logger"
	"server/pkg/security"
	"server/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SigninHandler(c fiber.Ctx) error {
	var input dto.AuthSinginRequest

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "Datos mal formateados"})
	}

	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	var user models.User
	tx := config.DB

	if input.Email != nil && *input.Email != "" {
		tx = tx.Where("email = ?", *input.Email)
	} else if input.DNI != nil && *input.DNI != "" {
		tx = tx.Where("dni = ?", *input.DNI)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "Debe proporcionar un email o DNI"})
	}

	if err := tx.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Error: "Credenciales inválidas"})
		}

		logger.Log.Error("Error consultando usuario:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Error: "Credenciales inválidas"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.AuthSinginResponse{UserID: user.ID.String(), Status: user.Status, IsDeleted: user.IsDeleted})
}
