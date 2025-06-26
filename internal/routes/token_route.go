package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterTokenRoutes(r fiber.Router) {
	tokens := r.Group("/tokens")
	tokens.Post("/generate", handlers.GenerateTokenHandler)
	tokens.Get("/validate", handlers.ValidateTokenHandler)
}
