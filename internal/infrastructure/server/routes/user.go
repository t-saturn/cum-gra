package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	"github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func UserRoutes(api fiber.Router) {
	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	structural_repository := postgres.NewStructuralPositionRepository()
	structural_service := services.NewStructuralPositionService(structural_repository)
	structural_handler := handlers.NewStructuralPositionHandler(structural_service)

	group := api.Group("/structural-positions")
	group.Post("/", structural_handler.Create())
	group.Get("/:id", structural_handler.GetByID())
	group.Patch("/:id", structural_handler.Update())

	/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
	organic_repository := postgres.NewOrganicUnitRepository()
	organic_service := services.NewOrganicUnitService(organic_repository)
	organic_handler := handlers.NewOrganicUnitHandler(organic_service)

	group = api.Group("/organic-units")
	group.Post("/", organic_handler.Create())
}
