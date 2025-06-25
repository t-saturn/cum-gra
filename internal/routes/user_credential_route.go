package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
)

func RegisterUserCredentialRoutes(router fiber.Router) {
	userGroup := router.Group("/credentials")

	userGroup.Post("/", handlers.InsertCredentialHandler)
	userGroup.Get("/", handlers.GetAllCredentialsHandler)
}
