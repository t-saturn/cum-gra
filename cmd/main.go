package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Ruta de prueba
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Auth Service Running 🚀")
	})

	log.Fatal(app.Listen(":3000"))
}
