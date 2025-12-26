package routes

import (
	handlers "server/internal/handlers/modules"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterModulesRoutes(router fiber.Router) {
	app := router.Group("/modules")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-clients"))
	
	protected.Get("/", handlers.GetModulesHandler)
	protected.Get("/stats", handlers.GetModulesStatsHandler)
	protected.Get("/all", handlers.GetAllModulesHandler)
	protected.Get("/:id", handlers.GetModuleByIDHandler)
	protected.Post("/", handlers.CreateModuleHandler)
	protected.Put("/:id", handlers.UpdateModuleHandler)
	protected.Delete("/:id", handlers.DeleteModuleHandler)
	protected.Patch("/:id/restore", handlers.RestoreModuleHandler)
}