package routes

import (
	"github.com/central-user-manager/internal/adapters/handlers"
	repo "github.com/central-user-manager/internal/adapters/repositories/postgres"
	"github.com/central-user-manager/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	// Structural Positions
	structuralRepo := repo.NewStructuralPositionRepository()
	structuralService := services.NewStructuralPositionService(structuralRepo)
	structuralHandler := handlers.NewStructuralPositionHandler(structuralService)

	api.Post("/positions", structuralHandler.Create)
}
