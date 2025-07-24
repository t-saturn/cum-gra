package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

// RegisterHealthRoutes define la ruta GET /health para monitoreo.
func RegisterHealthRoutes(router fiber.Router) {
	router.Get("/health", func(c fiber.Ctx) error {
		response := dto.HealthResponse{
			Status:  "ok",
			Message: "Auth service is healthy",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	router.Get("/device-info", handlers.DeviceInfo)
}
