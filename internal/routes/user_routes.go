package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/handlers"
)

// RegisterUserRoutes agrupa rutas relacionadas a usuarios.
func RegisterUserRoutes(app fiber.Router) {
	user := app.Group("/users")
	user.Post("/", handlers.CreateUser)
}
