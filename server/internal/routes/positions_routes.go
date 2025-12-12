package routes

import (
	"server/internal/handlers/positions"

	"github.com/gofiber/fiber/v3"
)

func RegisterPositionRoutes(router fiber.Router) {
	app := router.Group("/positions")

	app.Get("/", handlers.GetPositionsHandler)
	app.Get("/stats", handlers.GetPositionsStatsHandler)
}
