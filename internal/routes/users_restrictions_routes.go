package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRestrictionsRoutes(router fiber.Router) {
	app := router.Group("/users-restrictions")

	app.Get("/", handlers.GetUsersRestrictionsHandler)
	app.Get("/stats", handlers.GetUsersRestrictionsStatsHandler)
}
