package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
)

type TokenRequest struct {
	UserID         string            `json:"user_id"`
	ApplicationID  string            `json:"application_id"`
	ApplicationURL string            `json:"application_url"`
	DeviceInfo     models.DeviceInfo `json:"deviceInfo"`
}

func GenerateTokenHandler(c fiber.Ctx) error {
	var body TokenRequest
	if err := c.Bind().Body(&body); err != nil || body.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	tokenStr, err := services.GenerateAndStoreToken(
		body.UserID,
		body.ApplicationID,
		body.ApplicationURL,
		body.DeviceInfo,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenStr})
}
