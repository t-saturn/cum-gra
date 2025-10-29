package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/handlers"
	"github.com/t-saturn/auth-service-server/internal/services"
)

func RegisterKeyRoutes(router fiber.Router) {
	// Grupo correcto: "/.well-known"
	key := router.Group("/.well-known")

	// Inyección de dependencias
	jwksSvc := services.NewJWKSService()
	jwksH := handlers.NewJWKSHandler(jwksSvc)

	// ¡Con slash inicial! -> "/.well-known/jwks.json"
	key.Get("/jwks.json", jwksH.GetJWKS)
}
