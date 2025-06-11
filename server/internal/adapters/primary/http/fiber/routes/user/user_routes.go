package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/handlers"
)

func SetupUserRoutes(app fiber.Router) {
	user := app.Group("/users")
	user.Post("/", handlers.CreateUser)
}
