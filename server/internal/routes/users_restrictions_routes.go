package routes

import (
	"server/internal/handlers/users"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRestrictionsRoutes(router fiber.Router) {
	app := router.Group("/users-restrictions")

	app.Get("/", handlers.GetUsersRestrictionsHandler)
	app.Get("/stats", handlers.GetUsersRestrictionsStatsHandler)
}
