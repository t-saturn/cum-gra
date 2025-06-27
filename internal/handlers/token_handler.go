package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
)

type TokenHandler struct {
	service services.TokenService
}

func NewTokenHandler(service services.TokenService) *TokenHandler {
	return &TokenHandler{service: service}
}

func (h *TokenHandler) GenerateToken(c fiber.Ctx) error {
	userID := c.Query("user_id")
	appID := c.Query("application_id")
	if userID == "" || appID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id y application_id son requeridos"})
	}

	deviceInfo := models.DeviceInfo{
		UserAgent:   string(c.Request().Header.UserAgent()),
		IP:          c.IP(),
		BrowserName: c.Get("Sec-CH-UA-Platform"), // Mejora esto con parser si quieres más exactitud
		DeviceID:    c.Cookies("device_id"),
	}

	token, tokenStr, err := h.service.GenerateAndStoreToken(c.Context(), userID, appID, deviceInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token":       tokenStr,
		"expires_at":  token.ExpiresAt,
		"token_id":    token.TokenID,
		"status":      token.Status,
		"application": token.ApplicationID,
	})
}
