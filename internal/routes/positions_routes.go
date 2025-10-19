package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterStructuralPositionRoutes(router fiber.Router) {
	position := router.Group("/positions")

	position.Get("/", handlers.GetPositionsHandler)
}
