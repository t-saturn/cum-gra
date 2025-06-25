package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterAllRoutes(app *fiber.App) {
	api := app.Group("/api") // puedes ajustar el prefijo según tu versión

	RegisterUserCredentialRoutes(api)
	RegisterTokenRoutes(api)
	RegisterActiveTokenRoutes(api)
}
