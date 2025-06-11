package routes

import (
	"github.com/gofiber/fiber/v2"
	userRoutes "github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber/routes/user"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	userRoutes.SetupUserRoutes(api)
}
