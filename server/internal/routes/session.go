package routes

import (
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterSessionMeRoutes(app *fiber.App) {
	service := services.NewSessionMeService(config.DB)
	handler := handlers.NewSessionMeHandler(service)

	app.Post("/session/me", handler.SessionMe)
}
