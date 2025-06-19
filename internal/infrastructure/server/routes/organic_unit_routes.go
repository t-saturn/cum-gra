package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterOrganicUnitRoutes(api fiber.Router) {
	repo := repo.NewOrganicUnitRepository()
	service := services.NewOrganicUnitService(repo)
	handler := handlers.NewOrganicUnitHandler(service)

	route := api.Group("/organic_units")
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
