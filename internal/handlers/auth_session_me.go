package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Me maneja GET /auth/session/me: obtiene la sesión actual del usuario.
func (h *AuthHandler) Me(c fiber.Ctx) error {
	// 1. Extraer token de la cabecera Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "MISSING_TOKEN", "Falta cabecera Authorization")
	}
	// En formato "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_AUTH_HEADER", "Formato de Authorization inválido")
	}
	token := parts[1]

	// 2. Punto de integración pendiente: llamar a h.authService.GetCurrentSession(...)
	logger.Log.Info("Aquí implementamos el servicio GetCurrentSession con token=", token)

	// 3. Devolver placeholder con data = null
	return utils.JSONResponse[*dto.SessionMeResponseDTO](c, http.StatusOK, true, "Sesión actual obtenida", nil, nil)
}
