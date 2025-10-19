package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterOrganicUnitRoutes(router fiber.Router) {
	ou := router.Group("/units")

	ou.Get("/", handlers.GetOrganicUnitsHandler)
}
