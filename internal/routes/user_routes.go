package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

// RegisterUserRoutes agrupa rutas relacionadas a usuarios.
func RegisterUserRoutes(app fiber.Router) {
	user := app.Group("/users")
	user.Post("/", handlers.CreateUser)
}
