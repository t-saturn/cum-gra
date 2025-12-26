package routes

import (
	handlers "server/internal/handlers/organic-units"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterOrganicUnitsRoutes(router fiber.Router) {
	app := router.Group("/organic-units")
	app.Use(middlewares.KeycloakAuth())

	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))

	// Plantilla y carga masiva
	protected.Get("/template", handlers.DownloadOrganicUnitsTemplateHandler)
	protected.Post("/bulk", handlers.BulkCreateOrganicUnitsHandler)

	// CRUD
	protected.Get("/", handlers.GetOrganicUnitsHandler)
	protected.Get("/stats", handlers.GetOrganicUnitsStatsHandler)
	protected.Get("/all", handlers.GetAllOrganicUnitsHandler)
	protected.Get("/:id", handlers.GetOrganicUnitByIDHandler)
	protected.Post("/", handlers.CreateOrganicUnitHandler)
	protected.Put("/:id", handlers.UpdateOrganicUnitHandler)
	protected.Delete("/:id", handlers.DeleteOrganicUnitHandler)
	protected.Patch("/:id/restore", handlers.RestoreOrganicUnitHandler)
}