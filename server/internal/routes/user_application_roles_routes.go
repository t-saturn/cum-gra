package routes

import (
	handlers "server/internal/handlers/user-application-roles"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserApplicationRolesRoutes(router fiber.Router) {
	app := router.Group("/user-application-roles")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))
	
	protected.Get("/", handlers.GetUserApplicationRolesHandler)
	protected.Get("/stats", handlers.GetUserApplicationRolesStatsHandler)
	protected.Get("/:id", handlers.GetUserApplicationRoleByIDHandler)
	protected.Post("/", handlers.CreateUserApplicationRoleHandler)
	protected.Post("/bulk-assign-roles-to-user", handlers.BulkAssignRolesToUserHandler)
	protected.Post("/bulk-assign-role-to-users", handlers.BulkAssignRoleToUsersHandler)
	protected.Patch("/:id/revoke", handlers.RevokeUserApplicationRoleHandler)
	protected.Patch("/:id/restore", handlers.RestoreUserApplicationRoleHandler)
	protected.Delete("/:id", handlers.DeleteUserApplicationRoleHandler)
	protected.Patch("/:id/undelete", handlers.UndeleteUserApplicationRoleHandler)
}