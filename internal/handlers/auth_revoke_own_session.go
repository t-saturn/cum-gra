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

type revokeOwnSessionBody struct {
	Token        string  `json:"token" validate:"required"`
	SessionID    string  `json:"session_id" validate:"required"` // sesión OBJETIVO a revocar
	Reason       *string `json:"reason,omitempty"`
	RevokedByApp *string `json:"revoked_by_app,omitempty"`
}

// DELETE /auth/sessions  (body: token + session_id [+ reason, revoked_by_app])
func (h *AuthHandler) RevokeOwnSession(c fiber.Ctx) error {
	var body revokeOwnSessionBody

	// 1) Parse body
	if err := c.Bind().Body(&body); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}
	// 2) Validar requeridos
	if err := validator.Validate.Struct(&body); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	auth := dto.AuthRequestDTO{
		Token:     body.Token,
		SessionID: body.SessionID, // TARGET a revocar
	}
	meta := dto.RevokeOwnSessionQueryDTO{
		Reason:       body.Reason,
		RevokedByApp: body.RevokedByApp,
	}

	// 3) Delegar en service
	resp, err := h.authService.RevokeOwnSession(c, auth, meta)
	if err != nil {
		switch err {
		case services.ErrInvalidToken:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Token inválido o inactivo", "Token no válido")
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionInactive:
			return utils.JSONError(c, http.StatusForbidden, "SESSION_INACTIVE", "Sesión inactiva o revocada", "La sesión está inactiva o ha sido revocada")
		case services.ErrForbiddenRevoke:
			return utils.JSONError(c, http.StatusForbidden, "FORBIDDEN", "No puedes revocar una sesión que no es tuya", "operación no permitida")
		default:
			// Errores de repos
			switch err {
			case repositories.ErrUserDisabled, repositories.ErrUserDeleted:
				return utils.JSONError(c, http.StatusUnauthorized, "ACCOUNT_INACTIVE", "Cuenta inactiva o eliminada", err.Error())
			case repositories.ErrUserNotFound:
				return utils.JSONError(c, http.StatusNotFound, "USER_NOT_FOUND", "Usuario no encontrado", err.Error())
			}
			logger.Log.Errorf("Error en DELETE /auth/sessions: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error interno al revocar sesión", "Error desconocido")
		}
	}

	// 4) OK
	return utils.JSONResponse(c, http.StatusOK, true, "OK", resp, nil)
}
