package routes

import (
	"server/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterModulesRoutes(router fiber.Router) {
	app := router.Group("/modules")

	app.Get("/", handlers.GetModulesHandler)
	app.Get("/stats", handlers.GetModulesStatsHandler)
}
