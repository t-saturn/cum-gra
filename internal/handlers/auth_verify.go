package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// AuthHandler agrupa los handlers relacionados a autenticación.
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
func NewAuthHandler() *AuthHandler {
	pgDB := config.GetPostgresDB()
	mongoDB := config.MongoDB

	service := services.NewAuthService(pgDB, mongoDB)
	return &AuthHandler{
		authService: service,
	}
}

// VerifyCredentials maneja la verificación de credenciales.
func (h *AuthHandler) Verify(c fiber.Ctx) error {
	var input dto.AuthVerifyRequestDTO

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

	// 3. Verificar credenciales
	result, err := h.authService.VerifyCredentials(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", models.AuthStatusInvalid, "Credenciales no válidas")
		case services.ErrInactiveAccount:
			return utils.JSONError(c, http.StatusUnauthorized, "INACTIVE_ACCOUNT", models.SessionStatusInactive, "La cuenta está inactiva")
		default:
			logger.Log.Errorf("Error en autenticación: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "AUTH_FAILED", models.AuthStatusFailed, "Error desconocido")
		}
	}

	// 4. Respuesta
	return utils.JSONResponse(c, http.StatusOK, result.Success, result.Message, result.Data, nil)
}
