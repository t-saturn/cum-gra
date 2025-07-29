package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequestDTO

	// Bindear el cuerpo JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// Validación de campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	authService := services.NewAuthService(config.GetPostgresDB())
	result, err := authService.VerifyCredentials(input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", models.AuthStatusInvalid)
		case services.ErrInactiveAccount:
			return utils.JSONError(c, http.StatusUnauthorized, "INACTIVE_ACCOUNT", models.SessionStatusInactive)
		default:
			logger.Log.Errorf("Error en autenticación: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "AUTH_FAILED", models.AuthStatusFailed)
		}
	}

	// Respuesta exitosa
	return utils.JSON(c, http.StatusOK, dto.AuthVerifyResponseDTO{
		UserID:       result.UserID,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}
