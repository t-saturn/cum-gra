package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterApplicationsRoutes(router fiber.Router) {
	app := router.Group("/applications")

	app.Get("/", handlers.GetApplicationsHandler)
}
