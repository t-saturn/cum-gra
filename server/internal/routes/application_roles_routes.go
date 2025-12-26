package routes

import (
	"server/internal/handlers/roles"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterApplicationRolesRoutes(router fiber.Router) {
	app := router.Group("/application-roles")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-clients"))
	
	protected.Get("/", handlers.GetApplicationRolesHandler)
	protected.Get("/stats", handlers.GetApplicationRolesStatsHandler)
	protected.Get("/all", handlers.GetAllApplicationRolesHandler)
	protected.Get("/:id", handlers.GetApplicationRoleByIDHandler)
	protected.Post("/", handlers.CreateApplicationRoleHandler)
	protected.Put("/:id", handlers.UpdateApplicationRoleHandler)
	protected.Delete("/:id", handlers.DeleteApplicationRoleHandler)
	protected.Patch("/:id/restore", handlers.RestoreApplicationRoleHandler)
}