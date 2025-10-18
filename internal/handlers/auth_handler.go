package handlers

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"
	"central-user-manager/pkg/security"
	"central-user-manager/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// VerifyCredentialsHandler maneja la solicitud POST para verificar credenciales de autenticación y retorna el ID del usuario si son válidas.
func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequest

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

	var user models.User
	tx := config.DB

	// Buscar por email o DNI
	if *input.Email != "" {
		tx = tx.Where("email = ?", input.Email)
	} else if *input.DNI != "" {
		tx = tx.Where("dni = ?", input.DNI)
	}

	if err := tx.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error: "Credenciales inválidas",
			})
		}
		logger.Log.Error("Error consultando usuario:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Error interno del servidor",
		})
	}

	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "Credenciales inválidas",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.AuthVerifyResponse{
		UserID: user.ID.String(),
	})
}
