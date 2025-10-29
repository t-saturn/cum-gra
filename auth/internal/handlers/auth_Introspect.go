package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

type errMapEntry struct {
	target  error
	status  int
	code    string
	message string
	details string
}

var introspectErrMap = []errMapEntry{
	{services.ErrMissingFields, http.StatusBadRequest, "MISSING_FIELDS", "Datos incompletos en la solicitud", "Faltan session_id/access_token/refresh_token"},
	{services.ErrSessionNotFound, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión"},
	{services.ErrSessionInactive, http.StatusUnauthorized, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada"},
	{services.ErrMalformedAccessToken, http.StatusUnprocessableEntity, "MALFORMED_ACCESS_TOKEN", "El access token es inválido o mal formado", "Access token no válido"},
	{services.ErrMalformedRefreshToken, http.StatusUnprocessableEntity, "MALFORMED_REFRESH_TOKEN", "El refresh token es inválido o mal formado", "Refresh token no válido"},
	{services.ErrTokenRevoked, http.StatusUnauthorized, "TOKEN_REVOKED", "Token revocado", "El token fue revocado"},
	{services.ErrTokenSessionMismatch, http.StatusUnprocessableEntity, "TOKEN_SESSION_MISMATCH", "Inconsistencia entre token y sesión", "sid/sub no coinciden"},
	{services.ErrTokenNotFound, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido"},
}

func writeMappedError(c fiber.Ctx, err error) error {
	for _, e := range introspectErrMap {
		if errors.Is(err, e.target) {
			return utils.JSONError(c, e.status, e.code, e.message, e.details)
		}
	}

	logger.Log.Errorf("Error en introspection: %v", err)
	return utils.JSONError(c, http.StatusInternalServerError, "INTROSPECTION_ERROR", "Error interno al validar tokens", "Error desconocido")
}

// estructura auxiliar para parsear el body
type IntrospectBodyInput struct {
	SessionID    string `json:"session_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// POST /auth/introspect
// SOLO cookies (HttpOnly) como fuente:
//   - session_id
//   - access_token
//   - refresh_token
func (h *AuthHandler) Introspect(c fiber.Ctx) error {
	sessionID, access, refresh, readErr := readIntrospectInputsCookiesOnly(c)
	if readErr != nil {
		return utils.JSONError(c, http.StatusBadRequest, "MISSING_FIELDS", "Faltan credenciales para introspección (session_id, access_token, refresh_token)", readErr.Error())
	}

	data, err := h.authService.IntrospectSessionTokens(c, sessionID, access, refresh)
	if err != nil {
		return writeMappedError(c, err)
	}

	return utils.JSONResponse(c, http.StatusOK, true, "Sesión válida", data, nil)
}

func readIntrospectInputsCookiesOnly(c fiber.Ctx) (sessionID, access, refresh string, err error) {
	sessionID = c.Cookies("session_id")
	access = c.Cookies("access_token")
	refresh = c.Cookies("refresh_token")

	if sessionID == "" || access == "" || refresh == "" {
		return "", "", "", errors.New("session_id, access_token y refresh_token son requeridos (cookies)")
	}

	return sessionID, access, refresh, nil
}

// func readIntrospectInputsBodyOnly(c fiber.Ctx) (sessionID, access, refresh string, err error) {
// 	var body IntrospectBodyInput
// 	if bindErr := c.Bind().Body(&body); bindErr != nil {
// 		return "", "", "", bindErr
// 	}

// 	if body.SessionID == "" || body.AccessToken == "" || body.RefreshToken == "" {
// 		return "", "", "", errors.New("session_id, access_token y refresh_token son requeridos (json body)")
// 	}

// 	return body.SessionID, body.AccessToken, body.RefreshToken, nil
// }
