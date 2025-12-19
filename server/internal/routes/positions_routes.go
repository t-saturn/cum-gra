package routes

import (
	handlers "server/internal/handlers/positions"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterPositionRoutes(router fiber.Router) {
	app := router.Group("/positions")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))
	
	protected.Get("/", handlers.GetStructuralPositionsHandler)
	protected.Get("/stats", handlers.GetStructuralPositionsStatsHandler)
	protected.Get("/:id", handlers.GetStructuralPositionByIDHandler)
	protected.Post("/", handlers.CreateStructuralPositionHandler)
	protected.Put("/:id", handlers.UpdateStructuralPositionHandler)
	protected.Delete("/:id", handlers.DeleteStructuralPositionHandler)
	protected.Patch("/:id/restore", handlers.RestoreStructuralPositionHandler)
}