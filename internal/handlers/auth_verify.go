package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequest

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Datos mal formateados",
		})
	}

	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	authService := services.NewAuthService(config.GetPostgresDB())

	result, err := authService.VerifyCredentials(input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error: "Credenciales inválidas",
			})
		case services.ErrInactiveAccount:
			return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error: "Cuenta inactiva",
			})
		default:
			logger.Log.Errorf("Error en autenticación: %v", err)
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error: "Error interno del servidor",
			})
		}
	}

	// TODO: Crear sesión y logs en MongoDB

	return c.Status(http.StatusOK).JSON(dto.AuthVerifyResponse{
		UserID:       result.UserID,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}
