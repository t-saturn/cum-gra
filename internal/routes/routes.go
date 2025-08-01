// Package routes contiene las definiciones de rutas HTTP de la aplicación.
package routes

import (
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// RegisterRoutes configura y registra todas las rutas HTTP de la aplicación.
func RegisterRoutes(app *fiber.App, pgDB *gorm.DB, mongoDB *mongo.Database, version string, deps map[string]string) {
	// Ruta raíz
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "auth service API is running",
		})
	})

	// Registrar rutas de autenticación
	RegisterAuthRoutes(app) // Asumiendo que RegisterAuthRoutes no necesita dependencias adicionales

	// Registrar rutas de salud
	RegisterHealthRoutes(app, pgDB, mongoDB, version, deps)
}
