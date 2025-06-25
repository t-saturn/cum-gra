package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterAllRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/auth/logs", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
