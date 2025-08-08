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

// POST /auth/logout?logout_type=...&session_id=...
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	// 1. Leer query params
	input := dto.LogoutRequestDTO{
		LogoutType: c.Query("logout_type"),
		SessionID:  c.Query("session_id"),
	}

	// 2. Validar query
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Extraer refresh token desde Authorization: Bearer ...
	authz := c.Get("Authorization")
	var refreshToken string
	if strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		refreshToken = strings.TrimSpace(authz[len("Bearer "):])
	}

	// Si no hay refresh token, error
	if refreshToken == "" {
		return utils.JSONError(c, http.StatusUnauthorized, "NO_TOKEN", "No se encontró refresh token en Authorization", "Envíe Authorization: Bearer <refresh_token>")
	}

	// 4. Lógica de servicio
	data, err := h.authService.Logout(c, refreshToken, input.LogoutType, input.SessionID)
	if err != nil {
		switch err {
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_INACTIVE", "Sesión ya inactiva", "La sesión ya está inactiva")
		default:
			logger.Log.Errorf("Error en logout: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "LOGOUT_FAILED", "Error interno al cerrar sesión", "Error desconocido")
		}
	}

	// 5. Responder con éxito
	return utils.JSONResponse(c, http.StatusOK, true, "Logout exitoso", data, nil)
}
