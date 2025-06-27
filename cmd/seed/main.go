// cmd/seed/main.go
package main

import (
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/internal/adapters/repositories/postgres/seeds"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/config"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
	"github.com/t-saturn/central-user-manager/pkg/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("The .env file could not be loaded, using current environment")
	}

	logger.InitLogger()

	config.LoadConfig()
	database.Connect()

	logger.Log.Info("executing-seeders")

	if err := seeds.Run(); err != nil {
		logger.Log.Fatalf("error-when-running-seeders: %v", err)
	}

	logger.Log.Info("seeders-executed-correctly")
}
