package handlers

import (
	"server/internal/dto"
	"server/internal/services"

	"github.com/gofiber/fiber/v3"
)

type SessionMeHandler struct {
	service services.SessionMeService
}

func NewSessionMeHandler(s services.SessionMeService) *SessionMeHandler {
	return &SessionMeHandler{service: s}
}

// POST /session/me
func (h *SessionMeHandler) SessionMe(c fiber.Ctx) error {
	var req dto.SessionMeRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "payload inválido",
		})
	}

	resp, err := h.service.Execute(req.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "usuario no encontrado",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    resp,
	})
}
