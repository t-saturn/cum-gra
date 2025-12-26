package routes

import (
	handlers "server/internal/handlers/positions"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterPositionRoutes(router fiber.Router) {
	app := router.Group("/positions")
	app.Use(middlewares.KeycloakAuth())

	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))

	// Plantilla y carga masiva
	protected.Get("/template", handlers.DownloadPositionsTemplateHandler)
	protected.Post("/bulk", handlers.BulkCreatePositionsHandler)

	// CRUD
	protected.Get("/", handlers.GetStructuralPositionsHandler)
	protected.Get("/stats", handlers.GetStructuralPositionsStatsHandler)
	protected.Get("/all", handlers.GetAllPositionsHandler)
	protected.Get("/:id", handlers.GetStructuralPositionByIDHandler)
	protected.Post("/", handlers.CreateStructuralPositionHandler)
	protected.Put("/:id", handlers.UpdateStructuralPositionHandler)
	protected.Delete("/:id", handlers.DeleteStructuralPositionHandler)
	protected.Patch("/:id/restore", handlers.RestoreStructuralPositionHandler)
}