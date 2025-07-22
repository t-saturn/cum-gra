package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

// RegisterAuthRoutes agrupa las rutas relacionadas con autenticación
func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/verify", handlers.VerifyCredentialsHandler)

	token := router.Group("/token")
	token.Post("/validate", handlers.ValidateTokenHandler)
}
