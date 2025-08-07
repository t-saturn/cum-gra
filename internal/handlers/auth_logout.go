package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Logout maneja POST /auth/logout: cierra la sesión y revoca los tokens.
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	// 1. Leemos logout_type desde query, no desde JSON
	logoutType := c.Query("logout_type")
	if logoutType == "" {
		return utils.JSONError(c, http.StatusBadRequest,
			"BAD_FORMAT", "Logout type missing", "logout_type debe ir como ?logout_type=…")
	}

	// 2. Validamos logout_type
	input := dto.LogoutQueryDTO{LogoutType: logoutType}
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest,
			dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	// 3. Leemos los tokens de cookies
	accessToken := c.Cookies("access_token")
	refreshToken := c.Cookies("refresh_token")
	if accessToken == "" && refreshToken == "" {
		return utils.JSONError(c, http.StatusUnauthorized,
			"NO_TOKEN", "No se encontró ningún token en cookies", "Debe estar autenticado")
	}

	fmt.Print("AccessToken:", accessToken, " RefreshToken:", refreshToken)

	// 4. Pasamos tokens al servicio en lugar de session_id
	data, err := h.authService.Logout(c, accessToken, refreshToken, input.LogoutType)
	if err != nil {
		switch err {
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound,
				"SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusBadRequest,
				"SESSION_INACTIVE", "Sesión ya inactiva", "La sesión ya está inactiva")
		default:
			logger.Log.Errorf("Error en logout: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError,
				"LOGOUT_FAILED", "Error interno al cerrar sesión", "Error desconocido")
		}
	}

	// 5. Limpiamos las cookies para realmente desloguear al cliente
	c.ClearCookie("access_token")
	c.ClearCookie("refresh_token")

	// 6. Devolvemos el resultado
	return utils.JSONResponse(c, http.StatusOK, true, "Logout exitoso", data, nil)
}
