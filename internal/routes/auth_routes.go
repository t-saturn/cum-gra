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
	auth.Get("/logout", authHandler.Logout)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/introspect", authHandler.Introspect)

	session := auth.Group("/session")
	session.Get("/me", authHandler.Me)
	session.Get("/list", authHandler.ListSessions)
	session.Delete("/", authHandler.Revoke)
}
