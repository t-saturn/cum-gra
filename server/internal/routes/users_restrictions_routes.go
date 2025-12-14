package routes

import (
	handlers "server/internal/handlers/user-restrictions"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRestrictionsRoutes(router fiber.Router) {
	app := router.Group("/user-restrictions")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))
	
	protected.Get("/", handlers.GetUserModuleRestrictionsHandler)
	protected.Get("/stats", handlers.GetUserModuleRestrictionsStatsHandler)
	protected.Get("/:id", handlers.GetUserModuleRestrictionByIDHandler)
	protected.Post("/", handlers.CreateUserModuleRestrictionHandler)
	protected.Post("/bulk-create", handlers.BulkCreateUserModuleRestrictionsHandler)
	protected.Put("/:id", handlers.UpdateUserModuleRestrictionHandler)
	protected.Delete("/:id", handlers.DeleteUserModuleRestrictionHandler)
	protected.Patch("/:id/restore", handlers.RestoreUserModuleRestrictionHandler)
}