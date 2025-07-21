// Package middlewares contiene middlewares HTTP reutilizables como CORS, logging y autenticación.
package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// CORSMiddleware configura el middleware de CORS para permitir
// peticiones desde otros orígenes.
// Es necesario para que el frontend (por ejemplo en otro puerto o dominio)
// pueda comunicarse con esta API.
func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	})
}
