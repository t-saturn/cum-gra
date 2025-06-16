// internal/infrastructure/server/fiber.go
package server

import (
	"github.com/central-user-manager/internal/infrastructure/server/middleware"
	"github.com/central-user-manager/internal/infrastructure/server/routes"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Middlewares globales
	middleware.Setup(app)

	// Definir rutas
	routes.SetupRoutes(app)
}
