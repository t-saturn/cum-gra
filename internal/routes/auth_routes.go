package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router) {
	app := router.Group("/auth")
	app.Post("/signin", handlers.SigninHandler)
	app.Post("/role", handlers.AuthRoleHandler)
}
