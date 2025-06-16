// cmd/api/main.go
package main

import (
	"log"
	"os"

	"github.com/central-user-manager/internal/infrastructure/server"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	app := fiber.New()

	// Middleware, Rutas, Configuración
	server.Setup(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor escuchando en http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
