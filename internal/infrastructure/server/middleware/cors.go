package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func EnableCORS(app *fiber.App) {
	app.Use(cors.New())
}
