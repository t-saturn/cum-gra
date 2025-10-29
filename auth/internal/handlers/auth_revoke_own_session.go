package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// DELETE /auth/sessions
// Header: Authorization: Bearer <access_token>
// Query: session_id (req), reason (req), revoked_by_app (opt)
func (h *AuthHandler) RevokeOwnSession(c fiber.Ctx) error {
	var meta dto.RevokeOwnSessionQueryDTO

	// 1. Parse query
	if err := c.Bind().Query(&meta); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_QUERY", "Parámetros de consulta inválidos", "query no válido")
	}

	// 2. Validar query (session_id requerido, reason requerido/oneof)
	if err := validator.Validate.Struct(&meta); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Extraer access token del header Authorization
	authz := c.Get("Authorization")
	var accessToken string
	if strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		accessToken = strings.TrimSpace(authz[len("Bearer "):])
	}
	if accessToken == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "NO_TOKEN", "No se encontró access token en Authorization", "Envíe Authorization: Bearer <access_token>")
	}

	// 4. Delegar en service
	resp, err := h.authService.RevokeOwnSession(c, accessToken, meta)
	if err != nil {
		switch err {
		case services.ErrInvalidToken:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")
		case services.ErrForbiddenRevoke:
			return utils.JSONError(c, http.StatusForbidden, "FORBIDDEN", "No puedes revocar una sesión que no es tuya", "operación no permitida")
		default:
			// Errores de repos
			switch err {
			case repositories.ErrUserDisabled, repositories.ErrUserDeleted:
				return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", err.Error())
			case repositories.ErrUserNotFound:
				return utils.JSONError(c, http.StatusNotFound, "USER_NOT_FOUND", "Usuario no encontrado", err.Error())
			}
			logger.Log.Errorf("Error en DELETE /auth/sessions: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al revocar sesión", "Error desconocido")
		}
	}

	// 5. OK
	return utils.JSONResponse(c, http.StatusOK, true, "OK", resp, nil)
}
