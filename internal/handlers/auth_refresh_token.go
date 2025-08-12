package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Refresh maneja POST /auth/token/refresh: genera nuevos tokens usando un refresh token.
// Header: Authorization: Bearer <refresh_token>
// Query:  session_id (requerido)
// Body:   { "device_info": { ... } }
func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	var q dto.AuthRefreshQueryDTO
	var b dto.AuthRefreshResquestDTO

	// 1. Parse query (session_id)
	if err := c.Bind().Query(&q); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_QUERY", "Parámetros de consulta inválidos", "query no válido")
	}

	// 2. Validar query
	if err := validator.Validate.Struct(&q); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Parse body (device_info)
	if err := c.Bind().Body(&b); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}
	// (Opcional) si tienes validaciones para DeviceInfoDTO, actívalas aquí
	// if err := validator.Validate.Struct(&b); err != nil { ... }

	// 4. Extraer refresh token del header Authorization
	authz := c.Get("Authorization")
	var refreshToken string
	if strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		refreshToken = strings.TrimSpace(authz[len("Bearer "):])
	}
	if refreshToken == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "NO_TOKEN", "No se encontró refresh token en Authorization", "Envíe Authorization: Bearer <refresh_token>")
	}

	// 5. Llamar al servicio de refresh
	resp, err := h.authService.RefreshToken(c, refreshToken, q, b)
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

	// 6. Responder con los nuevos tokens
	return utils.JSONResponse(c, http.StatusOK, true, "Token refrescado exitosamente", resp, nil)
}
