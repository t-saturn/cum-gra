package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	RegisterStructuralPositionRoutes(api)

	// authRoute := app.Group("/auth")
	// RegisterAuthRoutes(authRoute)
}
