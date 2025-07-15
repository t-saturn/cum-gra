package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	UserRoutes(api)

	// authRoute := app.Group("/auth")
	// RegisterAuthRoutes(authRoute)
}
