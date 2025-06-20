package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterModuleRolePermissionRoutes(api fiber.Router) {
	r := repo.NewModulePermissionRepository()
	s := services.NewModulePermissionService(r)
	h := handlers.NewModulePermissionHandler(s)

	group := api.Group("/module-role-permissions")
	group.Post("/", h.Create)
	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
