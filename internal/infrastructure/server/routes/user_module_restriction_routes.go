package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterUserModuleRestrictionRoutes(api fiber.Router) {
	r := repo.NewUserModuleRestrictionRepository()
	s := services.NewUserModuleRestrictionService(r)
	h := handlers.NewUserModuleRestrictionHandler(s)

	group := api.Group("/user_module_restrictions")
	group.Post("/", h.Create)
	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Put("/:id", h.Update)
	group.Delete("/:id", h.Delete)
}
