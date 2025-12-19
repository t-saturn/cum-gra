package routes

import (
	handlers "server/internal/handlers/module-role-permissions"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterModuleRolePermissionsRoutes(router fiber.Router) {
	app := router.Group("/module-role-permissions")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-clients"))
	
	protected.Get("/", handlers.GetModuleRolePermissionsHandler)
	protected.Get("/stats", handlers.GetModuleRolePermissionsStatsHandler)
	protected.Get("/:id", handlers.GetModuleRolePermissionByIDHandler)
	protected.Post("/", handlers.CreateModuleRolePermissionHandler)
	protected.Post("/bulk-assign", handlers.BulkAssignPermissionsHandler)
	protected.Put("/:id", handlers.UpdateModuleRolePermissionHandler)
	protected.Delete("/:id", handlers.DeleteModuleRolePermissionHandler)
	protected.Patch("/:id/restore", handlers.RestoreModuleRolePermissionHandler)
}