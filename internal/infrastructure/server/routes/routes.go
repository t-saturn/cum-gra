package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func Ping(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Api ping
	api.Get("/ping", Ping)

	// Structural Positions
	structuralRepo := repo.NewStructuralPositionRepository()
	structuralService := services.NewStructuralPositionService(structuralRepo)
	structuralHandler := handlers.NewStructuralPositionHandler(structuralService)

	positions := api.Group("/structural_positions")
	positions.Post("/", structuralHandler.Create)
	positions.Get("/", structuralHandler.GetAll)
	positions.Get("/:id", structuralHandler.GetByID)
	positions.Put("/:id", structuralHandler.Update)
	positions.Delete("/:id", structuralHandler.Delete)

	unitRepo := repo.NewOrganicUnitRepository()
	unitService := services.NewOrganicUnitService(unitRepo)
	unitHandler := handlers.NewOrganicUnitHandler(unitService)

	units := api.Group("/organic_units")
	units.Post("/", unitHandler.Create)
	units.Get("/", unitHandler.GetAll)
	units.Get("/:id", unitHandler.GetByID)
	units.Put("/:id", unitHandler.Update)
	units.Delete("/:id", unitHandler.Delete)

	appRepo := repo.NewApplicationRepository()
	appService := services.NewApplicationService(appRepo)
	appHandler := handlers.NewApplicationHandler(appService)

	apps := api.Group("/applications")
	apps.Post("/", appHandler.Create)
	apps.Get("/", appHandler.GetAll)
	apps.Get("/:id", appHandler.GetByID)
	apps.Put("/:id", appHandler.Update)
	apps.Delete("/:id", appHandler.Delete)

}
