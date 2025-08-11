package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/services"
)

type JWKSHandler struct {
	svc *services.JWKSService
}

func NewJWKSHandler(svc *services.JWKSService) *JWKSHandler {
	return &JWKSHandler{svc: svc}
}

func (h *JWKSHandler) GetJWKS(c fiber.Ctx) error {
	set, maxAge, err := h.svc.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "jwks_build_failed",
		})
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.Set(fiber.HeaderCacheControl, fmt.Sprintf("public, max-age=%d", maxAge))
	return c.JSON(set) // JWKS puro: {"keys":[...]}
}
