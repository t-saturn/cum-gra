package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v3"
)

func Ping(c fiber.Ctx) error {
	return c.SendString("Pong!")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Api ping
	api.Get("/ping", Ping)

	// Structural Positions
	structuralRepo := repo.NewStructuralPositionRepository()
	structuralService := services.NewStructuralPositionService(structuralRepo)
	structuralHandler := handlers.NewStructuralPositionHandler(structuralService)

	api.Post("/positions", structuralHandler.Create)
}
