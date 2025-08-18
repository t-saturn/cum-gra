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

// Logout maneja GET /auth/logout?logout_type=...&session_id=...
// Función: Cierra una sesión activa validando el refresh_token enviado en el header Authorization.
// Requiere:
//   - Header: Authorization: Bearer <refresh_token>
//   - Query:  logout_type (opcional), session_id (opcional)
//
// Devuelve: JSON con la información de cierre de sesión o error tipificado.
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	// 1. Leer parámetros de query (logout_type, session_id)
	input := dto.LogoutRequestDTO{
		LogoutType: c.Query("logout_type"),
		SessionID:  c.Query("session_id"),
	}

	// 2. Validar parámetros de entrada
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Extraer refresh token desde el header Authorization
	authz := c.Get("Authorization")
	var refreshToken string
	if strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		refreshToken = strings.TrimSpace(authz[len("Bearer "):])
	}

	// 3.1 Validar existencia del refresh token
	if refreshToken == "" {
		return utils.JSONError(
			c,
			http.StatusUnauthorized,
			"NO_TOKEN",
			"No se encontró refresh token en Authorization",
			"Envíe Authorization: Bearer <refresh_token>",
		)
	}

	// 4. Llamar al servicio de logout → revoca la sesión y los tokens asociados
	data, err := h.authService.Logout(c, refreshToken, input.LogoutType, input.SessionID)
	if err != nil {
		switch err {
		// 4.1 Sesión no encontrada
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")

		// 4.2 Sesión ya estaba inactiva
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_INACTIVE", "Sesión ya inactiva", "La sesión ya está inactiva")

		// 4.3 Error inesperado → log y 500
		default:
			logger.Log.Errorf("Error en logout: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "LOGOUT_FAILED", "Error interno al cerrar sesión", "Error desconocido")
		}
	}

	// 5. Éxito → devolver respuesta JSON con datos de la sesión/token revocados
	return utils.JSONResponse(c, http.StatusOK, true, "Logout exitoso", data, nil)
}
