// internal/infrastructure/server/middleware/cors.go (opcional si vas a usar desde el principio)
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func EnableCORS(app *fiber.App) {
	app.Use(cors.New())
}
