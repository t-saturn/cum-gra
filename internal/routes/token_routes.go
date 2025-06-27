package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/handlers"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/internal/services"
)

func RegisterTokenRoutes(router fiber.Router, db config.DatabaseConfig) {
	auth := router.Group("/token")

	repo := repositories.NewTokenRepository(db.Mongo)
	svc := services.NewTokenService(repo)
	handler := handlers.NewTokenHandler(svc)

	auth.Get("/generate", handler.GenerateToken)
}
