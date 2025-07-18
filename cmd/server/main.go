package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/config"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/server"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"github.com/t-saturn/central-user-manager/pkg/validator"
)

func main() {
	_ = godotenv.Load()

	logger.InitLogger()

	logger.Log.Info("starting-server")

	config.LoadConfig()
	database.Connect()

	validator.InitValidator()

	app := fiber.New()
	server.Setup(app)

	port := config.GetConfig().SERVERPort

	logger.Log.Infof("server-listening-in http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("error-at-the-start-of-the-server: %v", err)
	}
}
