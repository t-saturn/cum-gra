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

// GET /auth/me?session_id=...
// Header: Authorization: Bearer <access_token>
func (h *AuthHandler) Me(c fiber.Ctx) error {
	var q dto.AuthMeQueryDTO
	// 1. Leer query params
	q.SessionID = c.Query("session_id")

	// 2. Validar query
	if err := validator.Validate.Struct(&q); err != nil {
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

	// 4. Delegar en el service
	data, err := h.authService.Me(c, accessToken, q)
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
			// Errores de repos (usuario deshabilitado/eliminado/no encontrado)
			switch err {
			case repositories.ErrUserDisabled, repositories.ErrUserDeleted:
				return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", err.Error())
			case repositories.ErrUserNotFound:
				return utils.JSONError(c, http.StatusNotFound, "USER_NOT_FOUND", "Usuario no encontrado", err.Error())
			}
			logger.Log.Errorf("Error en /auth/me: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al obtener la sesión actual", "Error desconocido")
		}
	}

	// 5. OK
	return utils.JSONResponse(c, http.StatusOK, true, "OK", data, nil)
}
