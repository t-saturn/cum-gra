package routes

import (
	"server/internal/handlers/roles"

	"github.com/gofiber/fiber/v3"
)

func RegisterRolesAppRoutes(router fiber.Router) {
	app := router.Group("/roles-app")

	app.Get("/", handlers.GetRolesAppHandler)
	app.Get("/stats", handlers.GetRolesAppStatsHandler)
}
