package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterAuthRoutes(r fiber.Router) {
	auth := r.Group("/auth")
	auth.Post("/login", handlers.LoginHandler)
}
