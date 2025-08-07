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
	var input dto.LogoutRequestDTO

	// 1. Solo parseamos session_id y logout_type
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest,
			"BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}

	// 2. Validamos session_id y logout_type
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest,
			dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	// 3. Leemos el token de la cookie
	token := c.Cookies("access_token")
	fmt.Println("Token de acceso:", token)
	if token == "" {
		return utils.JSONError(c, http.StatusUnauthorized,
			"NO_TOKEN", "Token no presente", "No se encontró cookie de token")
	}

	// 4. Ejecutamos el logout (service ya usa solo input.SessionID)
	data, err := h.authService.Logout(c, input)
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
