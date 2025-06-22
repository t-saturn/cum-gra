// routes/auth_routes.go
package routes

import (
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/external/crypto"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	"github.com/t-saturn/central-user-manager/internal/core/services"
)

func RegisterAuthRoutes(api fiber.Router) {
	hashService := crypto.NewBcryptService()
	repo := repo.NewAuthRepository()
	service := services.NewAuthService(repo, hashService)
	handler := handlers.NewAuthHandler(service)

	auth := api.Group("/auth")
	auth.Post("/verify", handler.Login)
}
