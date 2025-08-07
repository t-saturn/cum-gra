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
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}

	// 2. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Llamar al servicio de refresh
	resp, err := h.authService.RefreshToken(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidToken, services.ErrInvalidTokenType:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_REFRESH_TOKEN", "Refresh token inválido o inactivo", "Token no válido")
		case services.ErrTokenExpired:
			return utils.JSONError(c, http.StatusUnauthorized, "REFRESH_EXPIRED", "Refresh token expirado", "Token expirado")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión asociada no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionInactive, services.ErrSessionExpired:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o expirada", "La sesión está inactiva o ha expirado")
		default:
			logger.Log.Errorf("Error al refrescar token: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "REFRESH_FAILED", "Error interno al refrescar token", "Error desconocido")
		}
	}

	// 4. Responder con los nuevos tokens
	return utils.JSONResponse(c, http.StatusOK, true, "Token refrescado exitosamente", resp, nil)
}
