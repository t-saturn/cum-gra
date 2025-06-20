package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterPasswordHistoryRoutes(api fiber.Router) {
	r := repo.NewPasswordHistoryRepository()
	s := services.NewPasswordHistoryService(r)
	h := handlers.NewPasswordHistoryHandler(s)

	group := api.Group("/password_history")
	group.Post("/", h.Create)
	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
