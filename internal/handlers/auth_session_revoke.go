package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Revoke maneja DELETE /auth/sessions/{session_id}: revoca una sesión y sus tokens.
func (h *AuthHandler) Revoke(c fiber.Ctx) error {
	// 1. Leer path param
	sessionID := c.Params("session_id")
	if sessionID == "" {
		return utils.JSONError(c, http.StatusBadRequest, "MISSING_SESSION_ID", "Falta session_id en la ruta")
	}

	// 2. Parsear body
	var input dto.SessionRevokeRequestDTO
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// 3. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 4. Punto de integración pendiente: llamar a h.authService.RevokeSession(...)
	logger.Log.Info("Aquí implementamos el servicio RevokeSession para session_id=", sessionID)

	// 5. Responder placeholder
	return utils.JSONResponse[*dto.SessionRevokeResponseDTO](c, http.StatusOK, true, "Sesión revocada exitosamente", nil)
}
