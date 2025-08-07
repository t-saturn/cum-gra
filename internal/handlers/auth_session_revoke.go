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

// Revoke maneja DELETE /auth/sessions/{session_id}: revoca una sesión y sus tokens.
func (h *AuthHandler) Revoke(c fiber.Ctx) error {
	// 1. Leer path query
	sessionID := c.Query("session_id")
	if sessionID == "" {
		return utils.JSONError(c, http.StatusBadRequest, "MISSING_SESSION_ID", "Falta session_id en la ruta", "falta session_id")
	}

	// 2. Parsear body
	var input dto.SessionRevokeRequestDTO
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados", "cuerpo no válido")
	}

	// 3. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 4. Llamar al servicio RevokeSession
	resp, err := h.authService.RevokeSession(c, sessionID, input.Reason, input.UserID, input.RevokedByApp)
	if err != nil {
		switch err {
		case services.ErrSessionNotFound:
			return utils.JSONError(c, http.StatusNotFound, "SESSION_NOT_FOUND", "Sesión no encontrada", "No se pudo encontrar la sesión")
		case services.ErrSessionAlreadyRevoked:
			return utils.JSONError(c, http.StatusConflict, "SESSION_ALREADY_REVOKED", "La sesión ya está revocada", "La sesión ya ha sido revocada")
		default:
			logger.Log.Errorf("Error revocando sesión: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "REVOKE_FAILED", "No se pudo revocar la sesión", "Error desconocido")
		}
	}

	// 5. Responder con datos reales
	return utils.JSONResponse(c, http.StatusOK, true, "Sesión revocada exitosamente", resp, nil)
}
