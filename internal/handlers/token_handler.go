package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
)

type TokenRequest struct {
	UserID string `json:"user_id"`
}

func GenerateTokenHandler(c fiber.Ctx) error {
	var body TokenRequest
	if err := c.Bind().Body(&body); err != nil || body.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	token, jti, exp, err := config.GenerateJWT(body.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token generation failed"})
	}

	// Puedes guardar en base de datos aquí si lo deseas.

	return c.JSON(fiber.Map{
		"token": token,
		"jti":   jti,
		"exp":   exp.Format(time.RFC3339),
	})
}
