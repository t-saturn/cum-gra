package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/config"
)

func main() {
	config.InitLogger()
	config.Logger.Info("Iniciando servicio de autenticación...")

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		config.Logger.WithFields(map[string]interface{}{
			"path": c.Path(),
			"ip":   c.IP(),
		}).Info("Solicitud recibida")

		return c.SendString("Auth Service Running 🚀")
	})

	if err := app.Listen(":3000"); err != nil {
		config.Logger.WithError(err).Fatal("Fallo al iniciar el servidor")
		log.Fatal(err)
	}
}
