package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRestrictionsRoutes(router fiber.Router) {
	app := router.Group("/user-restrictions")

	app.Get("/", handlers.GetUserRestrictionsHandler)
	// app.Get("/stats", handlers.GetUserRestrictionsStatsHandler)
}
