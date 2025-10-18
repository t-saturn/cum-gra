// Package main inicia el servidor principal de la aplicación usando Fiber.
package main

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/middlewares"
	"central-user-manager/internal/routes"
	"central-user-manager/pkg/logger"
	"central-user-manager/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger.InitLogger()
	logger.Log.Info("Iniciando servidor...")

	config.LoadConfig()
	config.ConnectDB()

	if err := validator.InitValidator(); err != nil {
		logger.Log.Fatalf("Error al inicializar el validador: %v", err)
	}

	app := fiber.New()

	// Aplica logger personalizado
	app.Use(middlewares.CORSMiddleware())
	app.Use(middlewares.LoggerMiddleware())

	routes.RegisterRoutes(app)

	port := config.GetConfig().SERVERPort
	logger.Log.Infof("server-listening-in http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("error-at-the-start-of-the-server: %v", err)
	}
}
