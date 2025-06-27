package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/routes"
)

func main() {
	// Inicialización
	config.LoadEnv()
	config.InitLogger()
	config.InitMongoDB()
	config.InitJWT()

	config.Logger.Info("Starting Auth Service...")

	// Obtener configuración de base de datos
	db := config.GetDatabaseConfig()

	// Obtener puerto
	port := config.GetEnv("PORT", "3000")
	config.Logger.Infof("Auth Service escuchando en el puerto %s...", port)

	// Crear servidor Fiber
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		config.Logger.WithFields(map[string]interface{}{
			"path": c.Path(),
			"ip":   c.IP(),
		}).Info("Solicitud recibida en /")

		return c.SendString("Auth Service Running")
	})

	// Registrar rutas
	routes.RegisterAllRoutes(app, db)

	// Iniciar servidor
	if err := app.Listen(":" + port); err != nil {
		config.Logger.WithError(err).Fatal("Error al iniciar el servidor")
		log.Fatal(err)
	}
}
