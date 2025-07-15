package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	"github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func UserRoutes(api fiber.Router) {
	repo := postgres.NewStructuralPositionRepository()
	service := services.NewStructuralPositionService(repo)
	handler := handlers.NewStructuralPositionHandler(service)

	group := api.Group("/structural-positions")
	group.Post("/", handler.Create())
	group.Get("/:id", handler.GetByID())
	group.Patch("/:id", handler.Update())
}
