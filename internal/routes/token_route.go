package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterTokenRoutes(router fiber.Router) {
	group := router.Group("/tokens")
	group.Get("/generate", handlers.GenerateTokenHandler)
	group.Get("/validate", handlers.ValidateTokenHandler)
}
