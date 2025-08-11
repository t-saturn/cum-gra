package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Me maneja POST /auth/me: recibe { token, session_id } en el body.
func (h *AuthHandler) Me(c fiber.Ctx) error {
	var input dto.AuthRequestDTO

	// 1) Parsear JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}

	// 2) Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3) Delegar en el service
	data, err := h.authService.Me(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidToken:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionMismatch:
			return utils.JSONError(c, http.StatusBadRequest, "SESSION_MISMATCH", "Token no pertenece a la sesión proporcionada", "Token no válido")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")
		default:
			// Errores de repos (usuario deshabilitado/eliminado/no encontrado)
			switch err {
			case repositories.ErrUserDisabled, repositories.ErrUserDeleted:
				return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", err.Error())
			case repositories.ErrUserNotFound:
				return utils.JSONError(c, http.StatusNotFound, "USER_NOT_FOUND", "Usuario no encontrado", err.Error())
			}
			logger.Log.Errorf("Error en /auth/me: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al obtener la sesión actual", "Error desconocido")
		}
	}

	// 4) OK
	return utils.JSONResponse(c, http.StatusOK, true, "OK", data, nil)
}
