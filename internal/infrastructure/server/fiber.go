package server

import (
	"github.com/t-saturn/central-user-manager/internal/infrastructure/server/middleware"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/server/routes"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	// Middlewares globales
	middleware.Setup(app)

	// Definir rutas
	routes.SetupRoutes(app)
}
