package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterPositionRoutes(router fiber.Router) {
	position := router.Group("/positions")

	position.Get("/", handlers.GetPositionsHandler)
	position.Get("/stats", handlers.GetPositionsStatsHandler)
}
