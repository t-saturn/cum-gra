package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

// ValidateTokenHandler recibe un token y devuelve información si es válido
func ValidateTokenHandler(c fiber.Ctx) error {
	var req dto.TokenValidationRequest

	if err := c.Bind().Body(&req); err != nil {
		logger.Log.Warnf("Body inválido: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.TokenValidationResponse{
			Valid:   false,
			Message: "Body inválido",
		})
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.TokenValidationResponse{
			Valid:   false,
			Message: "Token requerido",
		})
	}

	result := security.ValidateToken(req.Token)

	if result.Code != 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.TokenValidationResponse{
			Valid:   false,
			Message: result.Message,
		})
	}

	// Calcular expires_in
	now := time.Now()
	expiresIn := int64(result.Claims.Expiry.Time().Sub(now).Seconds())
	if expiresIn < 0 {
		expiresIn = 0
	}

	return c.Status(fiber.StatusOK).JSON(dto.TokenValidationResponse{
		Valid:     true,
		Message:   "Token válido",
		Subject:   result.Claims.Subject,
		IssuedAt:  result.Claims.IssuedAt.Time().Format(time.RFC3339),
		ExpiresAt: result.Claims.Expiry.Time().Format(time.RFC3339),
		ExpiresIn: expiresIn,
	})
}
