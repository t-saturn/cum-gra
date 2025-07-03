package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	RegisterStructuralPositionRoutes(api)
	RegisterOrganicUnitRoutes(api)
	RegisterApplicationRoutes(api)
	RegisterModuleRoutes(api)
	RegisterUserRoutes(api)
	RegisterUserApplicationRoleRoutes(api)
	RegisterApplicationRoleRoutes(api)
	RegisterModuleRolePermissionRoutes(api)
	RegisterPasswordHistoryRoutes(api)
	RegisterUserModuleRestrictionRoutes(api)

	authRoute := app.Group("/auth")
	RegisterAuthRoutes(authRoute)
}
