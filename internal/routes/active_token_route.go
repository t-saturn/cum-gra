package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterActiveTokenRoutes(router fiber.Router) {
	group := router.Group("/active-tokens")
	group.Get("/", handlers.GetAllActiveTokensHandler)
}
