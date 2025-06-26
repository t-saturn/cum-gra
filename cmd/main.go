package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/routes"
)

func main() {
	config.LoadEnv()
	config.InitLogger()
	config.InitMongoDB()
	config.InitJWT()

	config.Logger.Info("Starting Auth Service...")

	port := config.GetEnv("PORT", "3000")
	config.Logger.Infof("Starting Auth Service on port %s...", port)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		config.Logger.WithFields(map[string]interface{}{
			"path": c.Path(),
			"ip":   c.IP(),
		}).Info("Application received in /")

		return c.SendString("Auth Service Running")
	})

	routes.RegisterAllRoutes(app)

	if err := app.Listen(":" + port); err != nil {
		config.Logger.WithError(err).Fatal("Error starting server")
		log.Fatal(err)
	}
}
