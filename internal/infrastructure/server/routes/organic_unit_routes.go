package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func RegisterOrganicUnitRoutes(api fiber.Router) {
	repository := repo.NewOrganicUnitRepository()
	service := services.NewOrganicUnitService(repository)
	handler := handlers.NewOrganicUnitHandler(service)

	route := api.Group("/organic_units")
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
