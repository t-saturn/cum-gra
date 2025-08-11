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

// ListSessions maneja POST /auth/sessions:
// Body: { token, session_id }, Query: filtros/paginación.
func (h *AuthHandler) ListSessions(c fiber.Ctx) error {
	var auth dto.AuthRequestDTO
	var q dto.ListSessionsQueryDTO

	// 1 Parse body (token + session_id)
	if err := c.Bind().Body(&auth); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}
	if err := validator.Validate.Struct(&auth); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 2 Parse query (filtros, paginación)
	if err := c.Bind().Query(&q); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_QUERY", "Parámetros de consulta inválidos", "query no válido")
	}

	// 3 Delegar en service
	data, err := h.authService.ListSessions(c, auth, q)
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
		default:
			logger.Log.Errorf("Error en /auth/sessions: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al listar sesiones", "Error desconocido")
		}
	}

	// 4 OK
	return utils.JSONResponse(c, http.StatusOK, true, "OK", data, nil)
}
