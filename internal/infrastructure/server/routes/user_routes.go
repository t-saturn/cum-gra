package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/security"
)

func RegisterUserRoutes(api fiber.Router) {
	hasher := security.NewArgon2Service()
	repo := repo.NewUserRepository()
	service := services.NewUserService(repo, hasher)
	handler := handlers.NewUserHandler(service)

	route := api.Group("/users")
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
