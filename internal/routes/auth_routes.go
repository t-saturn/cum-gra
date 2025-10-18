package routes

import (
	"central-user-manager/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/verify", handlers.VerifyCredentialsHandler)
}
