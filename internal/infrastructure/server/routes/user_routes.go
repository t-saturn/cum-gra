package routes

import (
	crypto "github.com/central-user-manager/internal/adapters/external/crypto"

	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(api fiber.Router) {
	hashService := crypto.NewBcryptService()
	repo := repo.NewUserRepository()
	service := services.NewUserService(repo, hashService)
	handler := handlers.NewUserHandler(service)

	route := api.Group("/users")
	route.Post("/", handler.Create)
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Delete("/:id", handler.Delete)
}
