package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Refresh maneja POST /auth/token/refresh: genera nuevos tokens usando un refresh token.
func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	var input dto.AuthRefreshRequestDTO

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

	// 3. Punto de integración pendiente: llamar a h.authService.RefreshToken(...)
	logger.Log.Info("Aquí implementamos el servicio RefreshToken")

	// 4. Devolver placeholder con data = null
	return utils.JSONResponse(c, http.StatusOK, true, "Token refrescado exitosamente", "")
}

// Validate maneja POST /auth/token/validate: valida un access token.
func (h *AuthHandler) Validate(c fiber.Ctx) error {
	var input dto.TokenValidationRequestDTO

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

	// 3. Llamar al servicio de validación
	data, err := h.authService.ValidateToken(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidToken:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada")
		case services.ErrSessionMismatch:
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_MISMATCH", "Token no pertenece a la sesión proporcionada")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada")
		default:
			logger.Log.Errorf("Error validando token: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "VALIDATION_ERROR", "Error interno al validar token")
		}
	}

	// 4. Responder con success, message y data
	return utils.JSONResponse(c, http.StatusOK, true, "Token válido", data)
}
