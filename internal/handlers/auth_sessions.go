package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// ListSessions maneja GET /auth/sessions: lista todas las sesiones del usuario.
func (h *AuthHandler) ListSessions(c fiber.Ctx) error {
	// 1. Extraer y validar JWT de Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "NO_AUTH_HEADER", "Encabezado de autorización requerido", "falta Authorization")
	}
	// Asumiendo que authHeader es "Bearer <token>", extraemos el token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader { // No se encontró "Bearer "
		return utils.JSONError(c, http.StatusUnauthorized, "INVALID_AUTH_HEADER", "Formato de encabezado inválido", "cuerpo no válido")
	}
	logger.Log.Info("Validando token con Authorization=", authHeader)

	// 2. Extraer claims para obtener userID y current sessionID
	// claims, err := security.ValidateToken(tokenString)
	// if err != nil {
	// 	return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido")
	// }
	// userID := claims.UserID // Ejemplo, ajusta según tu implementación

	// 3. Parsear query params en el DTO
	var q dto.ListSessionsQueryDTO
	if err := c.Bind().Query(&q); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_QUERY", "Parámetros de consulta inválidos", "cuerpo no válido")
	}

	// 4. Punto de integración pendiente: llamar a h.authService.ListSessions(...)
	logger.Log.Info("Llamando al servicio ListSessions con query=", q)

	// 5. Responder placeholder con data = null
	return utils.JSONResponse[*dto.ListSessionsResponseDTO](c, http.StatusOK, true, "Sesiones obtenidas", nil, nil)
}
