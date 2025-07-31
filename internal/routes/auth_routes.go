package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

// RegisterAuthRoutes agrupa las rutas relacionadas con autenticación
func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	authHandler := handlers.NewAuthHandler()

	auth.Post("/verify", authHandler.Verify)
	auth.Post("/login", authHandler.Login)
	// auth.Post("/validate", authHandler.ValidateToken) // Este método debes implementarlo de forma similar
}
