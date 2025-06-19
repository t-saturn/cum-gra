package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

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

	modRepo := repo.NewModuleRepository()
	modService := services.NewModuleService(modRepo)
	modHandler := handlers.NewModuleHandler(modService)

	modules := api.Group("/modules")
	modules.Post("/", modHandler.Create)
	modules.Get("/", modHandler.GetAll)
	modules.Get("/:id", modHandler.GetByID)
	modules.Put("/:id", modHandler.Update)
	modules.Delete("/:id", modHandler.Delete)
}
