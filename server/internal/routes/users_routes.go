package routes

import (
	"server/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/stats", handlers.GetUsersStatsHandler)

	user.Get("/", handlers.GetUsersHandler)
	user.Post("/", handlers.CreateUserHandler)
}
