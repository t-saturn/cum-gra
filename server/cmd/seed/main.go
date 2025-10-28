package main

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/database/seeds"
	"central-user-manager/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	logger.InitLogger()
	logger.Log.Info("Iniciando seeder...")

	config.LoadConfig()
	config.ConnectDB()

	logger.Log.Info("executing seeders")

	if err := seeds.Run(); err != nil {
		logger.Log.Fatalf("error when runing seeders: %v", err)
	}

	logger.Log.Info("seeders executed correctly")
}
