package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/handlers"
)

// RegisterAuthRoutes define las rutas relacionadas con autenticación dentro del grupo /auth.
func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/verify", handlers.VerifyCredentialsHandler)
}
