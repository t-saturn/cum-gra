package routes

import (
	"server/internal/handlers/auth"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router) {
	app := router.Group("/auth")
	app.Post("/role", handlers.AuthRoleHandler)
}
