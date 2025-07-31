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

func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequestDTO

	// 1 Parsear JSON
	if err := c.Bind().Body(&input); err != nil {
		return utils.JSONError(c, http.StatusBadRequest, "BAD_FORMAT", "Datos mal formateados")
	}

	// 2 Validación de campos
	if err := validator.Validate.Struct(&input); err != nil {
		return utils.JSON(c, http.StatusBadRequest, dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// 3 Construir servicio con Postgres y MongoDB
	pgDB := config.GetPostgresDB()
	mongoDB := config.MongoDB // *mongo.Database, inicializado en config.ConnectMongo
	authService := services.NewAuthService(pgDB, mongoDB)

	// 4 Verificar credenciales
	result, err := authService.VerifyCredentials(c, input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			return utils.JSONError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", models.AuthStatusInvalid)
		case services.ErrInactiveAccount:
			return utils.JSONError(c, http.StatusUnauthorized, "INACTIVE_ACCOUNT", models.SessionStatusInactive)
		default:
			logger.Log.Errorf("Error en autenticación: %v", err)
			return utils.JSONError(c, http.StatusInternalServerError, "AUTH_FAILED", models.AuthStatusFailed)
		}
	}

	// 5 Devolver resultado DTO directamente
	return utils.JSONResponse(c, http.StatusOK, result.Success, result.Message, result.Data)
}
