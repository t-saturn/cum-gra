package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
)

// Estructura del request JSON
type LoginRequest struct {
	Email          string                   `json:"email"`
	Password       string                   `json:"password"`
	ApplicationID  string                   `json:"application_id"`
	ApplicationURL string                   `json:"application_url"`
	DeviceInfo     models.SessionDeviceInfo `json:"device_info"`
}

func LoginHandler(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	response, err := services.LoginUser(req.Email, req.Password, req.ApplicationID, req.ApplicationURL, req.DeviceInfo)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response)
}
