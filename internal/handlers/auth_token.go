package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Validate maneja POST /auth/token/validate: valida un access token.
func (h *AuthHandler) Validate(c fiber.Ctx) error {
	var input dto.TokenValidationRequestDTO

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

	// 3. Punto de integración pendiente: llamar a h.authService.ValidateToken(...)
	logger.Log.Info("Aquí implementamos el servicio ValidateToken")

	// 4. Devolver placeholder con data = null
	return utils.JSONResponse[*dto.TokenValidationResponseDTO](c, http.StatusOK, true, "Token válido", nil)
}
