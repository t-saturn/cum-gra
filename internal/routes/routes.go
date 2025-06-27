package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
)

func RegisterAllRoutes(app *fiber.App, db config.DatabaseConfig) {
	api := app.Group("/api")
	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	RegisterTokenRoutes(api, db)
}
