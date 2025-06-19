package main

import (
	"github.com/central-user-manager/internal/infrastructure/config"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/central-user-manager/internal/infrastructure/server"
	"github.com/central-user-manager/pkg/logger"
	"github.com/central-user-manager/pkg/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger.InitLogger()

	logger.Log.Info("Iniciando servidor...")

	config.LoadConfig()
	database.Connect()

	validator.InitValidator()

	app := fiber.New()
	server.Setup(app)

	port := config.GetConfig().SERVERPort

	logger.Log.Infof("Servidor escuchando en http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
