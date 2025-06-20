package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func RegisterApplicationRoutes(api fiber.Router) {
	repo := repo.NewApplicationRepository()
	service := services.NewApplicationService(repo)
	handler := handlers.NewApplicationHandler(service)

	route := api.Group("/applications")
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
