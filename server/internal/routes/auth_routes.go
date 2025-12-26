package routes

import (
	handlers "server/internal/handlers/auth"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router) {
	app := router.Group("/auth")

	// Aplicar middleware de autenticación
	app.Use(middlewares.KeycloakAuthOnly())

	// Ruta protegida
	app.Post("/role", handlers.AuthRoleHandler)
}