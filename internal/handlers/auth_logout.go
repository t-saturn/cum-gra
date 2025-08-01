package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Logout maneja POST /auth/logout: cierra la sesión y revoca los tokens.
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	var input dto.LogoutRequestDTO

	// 1. Parsear JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// 2. Validar campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Punto de integración pendiente: llamar a h.authService.Logout(...)
	logger.Log.Info("Aquí implementamos el servicio Logout")

	// 4. Devolver placeholder con data = null
	return utils.JSONResponse[*dto.LogoutResponseDTO](c, http.StatusOK, true, "Logout exitoso", nil)
}
