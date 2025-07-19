package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/handlers"
)

func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/verify", handlers.VerifyCredentialsHandler)
}
