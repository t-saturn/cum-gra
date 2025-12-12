package routes

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Central User Manager API is running"})
	})

	api := app.Group("/api")

	RegisterAuthRoutes(app)
	RegisterUserRoutes(app)
	RegisterPositionRoutes(app)
	RegisterOrganicUnitRoutes(app)
	RegisterApplicationsRoutes(api)
	RegisterModulesRoutes(app)
	RegisterRolesAppRoutes(app)
	RegisterRolesAssignmentsRoutes(app)
	RegisterUserRestrictionsRoutes(app)
}
