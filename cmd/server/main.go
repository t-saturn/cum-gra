// Package main inicia el servidor principal de la aplicación usando Fiber.
package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/internal/middlewares"
	"github.com/t-saturn/central-user-manager/internal/routes"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"github.com/t-saturn/central-user-manager/pkg/validator"
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
