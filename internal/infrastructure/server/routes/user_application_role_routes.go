package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func RegisterUserApplicationRoleRoutes(api fiber.Router) {
	r := repo.NewUserApplicationRoleRepository()
	s := services.NewUserApplicationRoleService(r)
	h := handlers.NewUserApplicationRoleHandler(s)

	group := api.Group("/user-application-roles")
	group.Post("/", h.Create)
	group.Get("/", h.GetAll)
}
