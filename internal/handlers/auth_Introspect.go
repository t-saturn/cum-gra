package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// POST /auth/introspect
// Precedencia de entrada: Cookies > Headers > Body JSON
// Headers esperados desde el API Gateway:
//
//	Authorization: Bearer <access_token>
//	Refresh-Token: <refresh_token>
//	X-Session-Id: <session_id>
func (h *AuthHandler) Introspect(c fiber.Ctx) error {
	sessionID, access, refresh, readErr := readIntrospectInputs(c)
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
			// Puedes usar 401 (invalid) o 404; 401 es común para tokens desconocidos.
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

func readIntrospectInputs(c fiber.Ctx) (sessionID, access, refresh string, err error) {
	cookieSID := c.Cookies("session_id")
	cookieAT := c.Cookies("access_token")
	cookieRT := c.Cookies("refresh_token")

	headerSID := c.Get("X-Session-Id")
	headerAT := c.Get("Authorization")
	headerRT := c.Get("Refresh-Token")

	logger.Log.Infof("[introspect][inputs] cookies: sid=%q at.len=%d rt.len=%d", cookieSID, len(cookieAT), len(cookieRT))
	logger.Log.Infof("[introspect][inputs] headers: sid=%q authz.present=%t rt.len=%d",
		headerSID, headerAT != "", len(headerRT))

	// 1) Cookies
	sessionID = cookieSID
	access = cookieAT
	refresh = cookieRT

	// 2) Headers
	if sessionID == "" {
		sessionID = headerSID
	}
	if access == "" {
		if strings.HasPrefix(strings.ToLower(headerAT), "bearer ") {
			access = strings.TrimSpace(headerAT[len("Bearer "):])
		}
	}
	if refresh == "" {
		refresh = headerRT
	}

	// 3) Body (fallback solo si POST/PUT/PATCH)
	if sessionID == "" || access == "" || refresh == "" {
		var in dto.IntrospectRequestDTO
		if err := c.Bind().Body(&in); err == nil {
			if sessionID == "" && in.SessionID != nil {
				sessionID = *in.SessionID
			}
			if access == "" && in.AccessToken != nil {
				access = *in.AccessToken
			}
			if refresh == "" && in.RefreshToken != nil {
				refresh = *in.RefreshToken
			}
		}
	}

	logger.Log.Infof("[introspect][inputs] chosen: sid=%q at.len=%d rt.len=%d", sessionID, len(access), len(refresh))

	if sessionID == "" || access == "" || refresh == "" {
		return "", "", "", errors.New("session_id, access_token y refresh_token son requeridos")
	}
	return sessionID, access, refresh, nil
}
