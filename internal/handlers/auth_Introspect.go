package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// POST /auth/introspect
// Fuente ÚNICA de credenciales: Cookies HttpOnly
// Cookies esperadas:
//   - session_id
//   - access_token
//   - refresh_token
func (h *AuthHandler) Introspect(c fiber.Ctx) error {
	sessionID, access, refresh, readErr := readIntrospectInputsCookiesOnly(c)
	if readErr != nil {
		return utils.JSONError(c, http.StatusBadRequest, "MISSING_FIELDS",
			"Faltan credenciales para introspección (session_id, access_token, refresh_token)",
			readErr.Error(),
		)
	}

	// Lógica principal: validar sesión + ambos tokens
	data, err := h.authService.IntrospectSessionTokens(c, sessionID, access, refresh)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrMissingFields):
			return utils.JSONError(c, http.StatusBadRequest, "MISSING_FIELDS",
				"Datos incompletos en la solicitud", "Faltan session_id/access_token/refresh_token")

		case errors.Is(err, services.ErrSessionNotFound):
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND",
				"Sesión no encontrada", "No se pudo encontrar la sesión")

		case errors.Is(err, services.ErrSessionInactive):
			return utils.JSONError(c, http.StatusUnauthorized, "SESSION_INACTIVE",
				"Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")

		case errors.Is(err, services.ErrMalformedAccessToken):
			return utils.JSONError(c, http.StatusUnprocessableEntity, "MALFORMED_ACCESS_TOKEN",
				"El access token es inválido o está mal formado", "Access token no válido")

		case errors.Is(err, services.ErrMalformedRefreshToken):
			return utils.JSONError(c, http.StatusUnprocessableEntity, "MALFORMED_REFRESH_TOKEN",
				"El refresh token es inválido o está mal formado", "Refresh token no válido")

		case errors.Is(err, services.ErrTokenRevoked):
			return utils.JSONError(c, http.StatusUnauthorized, "TOKEN_REVOKED",
				"Token revocado", "El token fue revocado")

		case errors.Is(err, services.ErrTokenSessionMismatch):
			return utils.JSONError(c, http.StatusUnprocessableEntity, "TOKEN_SESSION_MISMATCH",
				"Inconsistencia entre token y sesión", "sid/sub no coinciden")

		case errors.Is(err, services.ErrTokenNotFound):
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN",
				"Token inválido o inactivo", "Token no válido")

		default:
			logger.Log.Errorf("Error en introspection: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTROSPECTION_ERROR",
				"Error interno al validar tokens", "Error desconocido")
		}
	}

	// OK: si algún token está expirado, viene reflejado en data.Tokens.<...>.Status = "expired"
	return utils.JSONResponse(c, http.StatusOK, true, "Sesión válida", data, nil)
}

// Solo cookies. Sin headers ni body.
func readIntrospectInputsCookiesOnly(c fiber.Ctx) (sessionID, access, refresh string, err error) {
	sessionID = c.Cookies("session_id")
	access = c.Cookies("access_token")
	refresh = c.Cookies("refresh_token")

	logger.Log.Infof("[introspect][inputs] cookies: sid=%q at.len=%d rt.len=%d",
		sessionID, len(access), len(refresh))

	if sessionID == "" || access == "" || refresh == "" {
		return "", "", "", errors.New("session_id, access_token y refresh_token son requeridos (cookies)")
	}
	return sessionID, access, refresh, nil
}
