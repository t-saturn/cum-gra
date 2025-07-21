package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

// RegisterAuthRoutes agrupa las rutas relacionadas con autenticación
func RegisterAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/request-token", handlers.AuthVerifyHandler)
	auth.Post("/validate", handlers.ValidateTokenHandler)
	// Aquí puedes agregar más rutas luego, como:
	// auth.Post("/refresh", handlers.RefreshHandler)
	// auth.Post("/logout", handlers.LogoutHandler)
}
