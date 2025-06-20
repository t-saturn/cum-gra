package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterApplicationRoleRoutes(api fiber.Router) {
	r := repo.NewApplicationRoleRepository()
	s := services.NewApplicationRoleService(r)
	h := handlers.NewApplicationRoleHandler(s)

	group := api.Group("/application_roles")
	group.Post("/", h.Create)
	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
