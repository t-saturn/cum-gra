package routes

import (
	handlers "server/internal/handlers/ubigeos"
	"server/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func RegisterUbigeosRoutes(router fiber.Router) {
	app := router.Group("/ubigeos")
	app.Use(middlewares.KeycloakAuth())

	protected := app.Group("")
	protected.Use(middlewares.RequireResourceRole("realm-management", "manage-users"))

	// Endpoints para selects
	protected.Get("/departments", handlers.GetDepartmentsHandler)
	protected.Get("/provinces", handlers.GetProvincesByDepartmentHandler)
	protected.Get("/districts", handlers.GetDistrictsByProvinceHandler)

	// Plantilla y carga masiva
	protected.Get("/template", handlers.DownloadUbigeosTemplateHandler)
	protected.Post("/bulk", handlers.BulkCreateUbigeosHandler)

	// CRUD
	protected.Get("/", handlers.GetUbigeosHandler)
	protected.Get("/stats", handlers.GetUbigeosStatsHandler)
	protected.Get("/:id", handlers.GetUbigeoByIDHandler)
	protected.Post("/", handlers.CreateUbigeoHandler)
	protected.Put("/:id", handlers.UpdateUbigeoHandler)
	protected.Delete("/:id", handlers.DeleteUbigeoHandler)
}