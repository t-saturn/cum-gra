package routes

import (
	handlers "server/internal/handlers/users"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(router fiber.Router) {
	app := router.Group("/users")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))
	
	protected.Get("/", handlers.GetUsersHandler)
	protected.Get("/stats", handlers.GetUsersStatsHandler)
	protected.Get("/:id", handlers.GetUserByIDHandler)
	protected.Post("/", handlers.CreateUserHandler)
	protected.Put("/:id", handlers.UpdateUserHandler)
	protected.Delete("/:id", handlers.DeleteUserHandler)
	protected.Patch("/:id/restore", handlers.RestoreUserHandler)
}