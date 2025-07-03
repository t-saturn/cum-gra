package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/central-user-manager/internal/adapters/handlers"
	repo "github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/t-saturn/central-user-manager/internal/core/services"
	"github.com/t-saturn/central-user-manager/internal/shared/security"
)

func RegisterAuthRoutes(api fiber.Router) {
	authRepo := repo.NewAuthRepository()
	hasher := security.NewArgon2Service()
	authService := services.NewAuthService(authRepo, hasher)
	handler := handlers.NewAuthHandler(authService)

	auth := api.Group("/")
	auth.Post("/validate-credentials", handler.Login)
}
