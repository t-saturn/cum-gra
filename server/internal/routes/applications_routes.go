package routes

import (
	"server/internal/handlers/applications"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterApplicationsRoutes(router fiber.Router) {
	app := router.Group("/applications")

	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-clients"))

	protected.Get("/", handlers.GetApplicationsHandler)
	protected.Get("/stats", handlers.GetApplicationsStatsHandler)
	protected.Get("/:id", handlers.GetApplicationByIDHandler)
	protected.Post("/", handlers.CreateApplicationHandler)
	protected.Put("/:id", handlers.UpdateApplicationHandler)
	protected.Delete("/:id", handlers.DeleteApplicationHandler)
	protected.Patch("/:id/restore", handlers.RestoreApplicationHandler)
}
