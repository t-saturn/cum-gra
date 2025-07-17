package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/pkg/logger"
)

func main() {
	_ = godotenv.Load()

	logger.InitLogger()
	logger.Log.Info("Iniciando servidor...")

	config.LoadConfig()
	config.ConnectDB()

	port := config.GetConfig().SERVERPort
	app := fiber.New()

	logger.Log.Infof("server-listening-in http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("error-at-the-start-of-the-server: %v", err)
	}
}
