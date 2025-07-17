package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	// Ruta raíz de prueba
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Central User Manager API is running",
		})
	})

	// Aquí irán los grupos: users, auth, tokens, etc.
	// e.g. RegisterUserRoutes(app)
}
