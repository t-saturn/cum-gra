// Package routes contiene las definiciones de rutas HTTP de la aplicación.
package routes

import (
	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes configura y registra todas las rutas HTTP de la aplicación.
func RegisterRoutes(app *fiber.App) {
	// Ruta raíz
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "auth service API is running",
		})
	})

	RegisterAuthRoutes(app)
	RegisterKeyRoutes(app)
}
