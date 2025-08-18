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

// Login maneja POST /auth/login: flujo completo de login.
// Recibe: JSON en el body con credenciales y parámetros de login.
// Devuelve: JSON con información de sesión y tokens, o errores tipificados.
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var input dto.AuthLoginRequestDTO

	// 1. Parsear el body como JSON → convertir datos de entrada al DTO correspondiente.
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// 2. Validar campos del DTO → usar validator para aplicar reglas de negocio/estructura.
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3. Llamar al servicio de login → delegar la lógica de autenticación al AuthService.
	result, err := h.authService.Login(c, input)
	if err != nil {
		switch err {
		// 3.1 Credenciales inválidas → devolver 401 Unauthorized con código tipificado.
		case services.ErrInvalidCredentials:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", models.AuthStatusInvalid, "Credenciales inválidas")

		// 3.2 Cuenta inactiva → devolver 403 Forbidden.
		case services.ErrInactiveAccount:
			return utils.JSONError(c, http.StatusForbidden, "ACCOUNT_INACTIVE", models.SessionStatusInactive, "Cuenta inactiva")

		// 3.3 Cualquier otro error inesperado → loggear y devolver 500.
		default:
			logger.Log.Errorf("Error en login: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "LOGIN_FAILED", "Error interno al realizar login", "Error desconocido")
		}
	}

	// 4. Éxito → devolver respuesta en JSON con los datos principales:
	//    - user_id: identificador del usuario autenticado
	//    - session: información de la sesión creada
	//    - tokens: access/refresh tokens generados
	//    - attempt_id: identificador del intento de login (útil para auditoría/logs)
	return utils.JSONResponse(c, http.StatusOK, true, "Login exitoso", fiber.Map{
		"user_id":    result.UserID,
		"session":    result.Session,
		"tokens":     result.Tokens,
		"attempt_id": result.AttemptID,
	}, nil)
}
