package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterRolesAppRoutes(router fiber.Router) {
	app := router.Group("/roles-app")

	app.Get("/", handlers.GetRolesAppHandler)
}
