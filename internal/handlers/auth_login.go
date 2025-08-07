package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Login maneja POST /auth/login: flujo completo de login.
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var input dto.AuthLoginRequestDTO

	// 1. Parsear JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// 2. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Llamar al servicio de login
	result, err := h.authService.Login(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", models.AuthStatusInvalid)
		case services.ErrInactiveAccount:
			return utils.JSONError(c, http.StatusForbidden, "ACCOUNT_INACTIVE", models.SessionStatusInactive)
		default:
			logger.Log.Errorf("Error en login: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "LOGIN_FAILED", "Error interno al realizar login")
		}
	}

	// 4. Devolver respuesta con datos del login
	return utils.JSONResponse(c, http.StatusOK, true, "Login exitoso", result, nil)
}
