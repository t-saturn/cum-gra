package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(app fiber.Router) {
	user := app.Group("/users")

	user.Get("/", handlers.GetUsersHandler)
	user.Post("/", handlers.CreateUserHandler)
}
