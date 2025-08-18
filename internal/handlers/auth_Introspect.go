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

// Introspect maneja GET /auth/introspect: valida un access/refresh token.
// Header: Authorization: Bearer <token>
// Query:  session_id (requerido)
func (h *AuthHandler) Introspect(c fiber.Ctx) error {
	var q dto.IntrospectQueryDTO

	// 1. Parsear query
	if err := c.Bind().Query(&q); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_QUERY", "Parámetros de consulta inválidos", "query no válido")
	}
	// 2. Validar query
	if err := validator.Validate.Struct(&q); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Extraer token del header Authorization
	authz := c.Get("Authorization")
	var rawToken string
	if strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		rawToken = strings.TrimSpace(authz[len("Bearer "):])
	}
	if rawToken == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "NO_TOKEN", "No se encontró token en Authorization", "Envíe Authorization: Bearer <token>")
	}

	// 4. Llamar al servicio
	data, err := h.authService.Introspect(c, rawToken, q)
	if err != nil {
		switch err {
		case services.ErrInvalidToken:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionMismatch:
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_MISMATCH", "Token no pertenece a la sesión proporcionada", "Token no válido")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")
		case services.ErrTokenExpired:
			return utils.JSONError(c, http.StatusUnauthorized, "TOKEN_EXPIRED", "El token ha expirado", "Token expirado")
		default:
			logger.Log.Errorf("Error validando token: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "VALIDATION_ERROR", "Error interno al validar token", "Error desconocido")
		}
	}

	// 5. OK
	return utils.JSONResponse(c, http.StatusOK, true, "Token válido", data, nil)
}
