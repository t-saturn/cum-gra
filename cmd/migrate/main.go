package main

import (
	"central-user-manager/internal/config"
	"central-user-manager/internal/database/migrations"
	"central-user-manager/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()

	_ = godotenv.Load()
	config.LoadConfig()
	config.ConnectDB()

	migrations.HandleMigration()
}
