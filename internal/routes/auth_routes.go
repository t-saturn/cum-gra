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
	auth.Post("/logout", authHandler.Logout)

	token := auth.Group("/token")
	token.Post("/validate", authHandler.Validate)
	token.Post("/refresh", authHandler.Refresh)

	session := auth.Group("/session")
	session.Get("/me", authHandler.Me)
}
