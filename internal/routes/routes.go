package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterAllRoutes(app *fiber.App) {
	api := app.Group("/api")

	RegisterUserCredentialRoutes(api)
	RegisterTokenRoutes(api)
	RegisterActiveTokenRoutes(api)
	RegisterInvalidTokenRoutes(api)
}
