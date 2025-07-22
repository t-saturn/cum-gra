package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/validator"
	"gorm.io/gorm"
)

// VerifyCredentialsHandler maneja la solicitud POST para verificar credenciales y generar tokens.
func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequest

	// Validar formato del body
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Datos mal formateados",
		})
	}

	// Validar campos requeridos
	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	// Obtener conexión a PostgreSQL
	db := config.GetPostgresDB()

	var user models.User
	tx := db.Model(&models.User{})

	// Buscar por email o DNI
	if input.Email != nil && *input.Email != "" {
		tx = tx.Where("email = ?", input.Email)
	} else if input.DNI != nil && *input.DNI != "" {
		tx = tx.Where("dni = ?", input.DNI)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Debe proporcionar un email o DNI",
		})
	}

	// Ejecutar consulta
	if err := tx.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error: "Credenciales inválidas",
			})
		}
		logger.Log.Errorf("Error consultando usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Error interno del servidor",
		})
	}

	// Verificar contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "Credenciales inválidas",
		})
	}

	// Generar tokens
	accessToken, err := security.GenerateToken(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "No se pudo generar el token de acceso",
		})
	}

	refreshToken, err := security.GenerateToken(user.ID.String()) // Idealmente, generar con expiración diferente
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "No se pudo generar el token de refresco",
		})
	}

	// TODO: Crear sesión en MongoDB (puedes agregar aquí la lógica cuando tengas el repositorio implementado)

	// Devolver respuesta con user_id y tokens
	return c.Status(fiber.StatusOK).JSON(dto.AuthVerifyResponse{
		UserID:       user.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
