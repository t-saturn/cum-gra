package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterInvalidTokenRoutes(router fiber.Router) {
	group := router.Group("/invalid-tokens")
	group.Post("/", handlers.CreateInvalidTokenHandler)
	group.Get("/", handlers.GetAllInvalidTokensHandler)
}
