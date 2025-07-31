package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// Login maneja la ruta POST /auth/login.
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var input dto.AuthLoginRequestDTO

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

	// 3. Punto de integración pendiente: llamar al servicio Login
	fmt.Println("Aquí llamamos al servicio Login (pendiente de implementación)")

	// 4. Devolver respuesta con data = null
	return utils.JSONResponse[*dto.AuthLoginResponseDTO](c, http.StatusOK, true, "Login exitoso", nil)
}
