package routes

import (
	handlers "server/internal/handlers/ubigeos"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterUbigeosRoutes(router fiber.Router) {
	app := router.Group("/ubigeos")
	
	// Aplicar KeycloakAuth a todas las rutas
	app.Use(middlewares.KeycloakAuth())

	// Rutas protegidas
	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))
	
	protected.Get("/", handlers.GetUbigeosHandler)
	protected.Get("/stats", handlers.GetUbigeosStatsHandler)
	protected.Get("/:id", handlers.GetUbigeoByIDHandler)
	protected.Post("/", handlers.CreateUbigeoHandler)
	protected.Put("/:id", handlers.UpdateUbigeoHandler)
	protected.Delete("/:id", handlers.DeleteUbigeoHandler)
}