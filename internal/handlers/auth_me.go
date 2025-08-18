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

// Me maneja GET /auth/me?session_id=...
// Función: Retorna la información de la sesión y el usuario autenticado.
// Requiere:
//   - Header: Authorization: Bearer <access_token>
//   - Query:  session_id (requerido)
//
// Devuelve: JSON con los datos del usuario/sesión o un error tipificado.
func (h *AuthHandler) Me(c fiber.Ctx) error {
	var q dto.AuthMeQueryDTO

	// 1. Leer parámetros de query (session_id)
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
	// 3.1 Validar que exista el token
	if accessToken == "" {
		return utils.JSONError(
			c,
			http.StatusUnauthorized,
			"NO_TOKEN",
			"No se encontró access token en Authorization",
			"Envíe Authorization: Bearer <access_token>",
		)
	}

	// 4. Delegar en el servicio → validar token + sesión y obtener datos de usuario
	data, err := h.authService.Me(c, accessToken, q)

	// 5. Mapeo de errores conocidos a respuestas tipificadas
	errorMap := map[error]func() error{
		services.ErrInvalidToken: func() error {
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido")
		},
		services.ErrSessionNotFound: func() error {
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		},
		services.ErrSessionMismatch: func() error {
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_MISMATCH", "Token no pertenece a la sesión proporcionada", "Token no válido")
		},
		services.ErrSessionInactive: func() error {
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")
		},
		repositories.ErrUserDisabled: func() error {
			return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", repositories.ErrUserDisabled.Error())
		},
		repositories.ErrUserDeleted: func() error {
			return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", repositories.ErrUserDeleted.Error())
		},
		repositories.ErrUserNotFound: func() error {
			return utils.JSONError(c, http.StatusNotFound, "USER_NOT_FOUND", "Usuario no encontrado", repositories.ErrUserNotFound.Error())
		},
	}

	// 6. Manejar errores devueltos por el servicio
	if err != nil {
		if handler, ok := errorMap[err]; ok {
			return handler()
		}
		// 6.1 Error desconocido → log y devolver 500
		logger.Log.Errorf("Error en /auth/me: %v", err)
		return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al obtener la sesión actual", "Error desconocido")
	}

	// 7. Éxito → devolver información de usuario/sesión
	return utils.JSONResponse(c, http.StatusOK, true, "OK", data, nil)
}
