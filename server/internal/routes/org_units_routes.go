package routes

import (
	"server/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterOrganicUnitRoutes(router fiber.Router) {
	app := router.Group("/units")

	app.Get("/", handlers.GetOrganicUnitsHandler)
	app.Get("/stats", handlers.GetOrganicUnitsStatsHandler)
}
