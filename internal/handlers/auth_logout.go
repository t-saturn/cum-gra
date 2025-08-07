package handlers

import (
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

	// 1. Parsear JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}

	// 2. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Llamar al servicio de logout
	data, err := h.authService.Logout(c, input)
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

	// 4. Devolver respuesta
	return utils.JSONResponse(c, http.StatusOK, true, "Logout exitoso", data, nil)
}
