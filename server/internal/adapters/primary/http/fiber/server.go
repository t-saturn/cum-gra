package fiberadapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/middleware"
)

func NewFiberServer() *fiber.App {
	app := fiber.New()

	// Middlewares
	app.Use(middleware.LoggerMiddleware())

	// Rutas básicas
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}
