package handlers

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

type UserData struct {
	ID               uuid.UUID  `json:"id"`
	Email            string     `json:"email"`
	PasswordHash     string     `json:"-"`
	DNI              string     `json:"dni"`
	EmailVerified    bool       `json:"email_verified"`
	TwoFactorEnabled bool       `json:"two_factor_enabled"`
	Status           string     `json:"status"`
	IsDeleted        bool       `json:"is_deleted"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DeletedBy        *uuid.UUID `json:"deleted_by"`
}

// VerifyCredentialsHandler maneja la solicitud POST para verificar credenciales y generar tokens.
func VerifyCredentialsHandler(c fiber.Ctx) error {
	var input dto.AuthVerifyRequest

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Datos mal formateados",
		})
	}

	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Errors: validator.FormatValidationError(err),
		})
	}

	db := config.GetPostgresDB()

	// Construir la consulta SQL dinámicamente
	var query string
	var args []interface{}

	if input.Email != nil && *input.Email != "" {
		query = `SELECT id, email, password_hash, dni, email_verified, two_factor_enabled,
				status, is_deleted, deleted_at, deleted_by
				FROM users WHERE email = ? AND is_deleted = false LIMIT 1`
		args = append(args, *input.Email)
	} else if input.DNI != nil && *input.DNI != "" {
		query = `SELECT id, email, password_hash, dni, email_verified, two_factor_enabled,
				status, is_deleted, deleted_at, deleted_by
				FROM users WHERE dni = ? AND is_deleted = false LIMIT 1`
		args = append(args, *input.DNI)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Debe proporcionar un email o DNI",
		})
	}

	// Ejecutar la consulta
	var userData UserData
	var deletedAtStr sql.NullString
	var deletedByStr sql.NullString

	row := db.Raw(query, args...).Row()
	err := row.Scan(
		&userData.ID,
		&userData.Email,
		&userData.PasswordHash,
		&userData.DNI,
		&userData.EmailVerified,
		&userData.TwoFactorEnabled,
		&userData.Status,
		&userData.IsDeleted,
		&deletedAtStr,
		&deletedByStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error: "Credenciales inválidas",
			})
		}
		logger.Log.Errorf("Error consultando usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Error interno del servidor",
		})
	}

	// Procesar campos nullable
	if deletedAtStr.Valid {
		deletedAt, _ := time.Parse(time.RFC3339, deletedAtStr.String)
		userData.DeletedAt = &deletedAt
	}
	if deletedByStr.Valid {
		deletedBy, _ := uuid.Parse(deletedByStr.String)
		userData.DeletedBy = &deletedBy
	}

	// Verificar si el usuario está activo
	if userData.Status != "active" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "Cuenta inactiva",
		})
	}

	// Verificar la contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, userData.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "Credenciales inválidas",
		})
	}

	// Generar tokens
	accessToken, err := security.GenerateToken(userData.ID.String(), 15*time.Minute)
	if err != nil {
		logger.Log.Errorf("Error generando access token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "No se pudo generar el token de acceso",
		})
	}

	refreshToken, err := security.GenerateToken(userData.ID.String(), 7*24*time.Hour)
	if err != nil {
		logger.Log.Errorf("Error generando refresh token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "No se pudo generar el token de refresco",
		})
	}

	// TODO: Crear sesión en MongoDB
	// TODO: Registrar AuthAttempt exitoso
	// TODO: Crear logs de tokens en MongoDB

	logger.Log.Infof("Usuario autenticado exitosamente: %s", userData.ID.String())

	// Devolver respuesta con user_id y tokens
	return c.Status(fiber.StatusOK).JSON(dto.AuthVerifyResponse{
		UserID:       userData.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
