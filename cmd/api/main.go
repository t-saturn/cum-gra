package main

import (
	"os"

	"github.com/central-user-manager/internal/infrastructure/server"
	"github.com/central-user-manager/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar .env
	_ = godotenv.Load()

	// Inicializar logger
	logger.InitLogger()
	logger.Log.Info("Iniciando servidor...")

	app := fiber.New()

	server.Setup(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log.Infof("Servidor escuchando en http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
