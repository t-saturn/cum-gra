// Package main ejecuta el proceso de seed para insertar datos iniciales en la base de datos.
package main

import (
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/internal/database/seeds"
	"github.com/t-saturn/central-user-manager/pkg/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("The .env file could not be loaded, using current environment")
	}

	logger.InitLogger()

	config.LoadConfig()
	config.ConnectDB()

	logger.Log.Info("executing seeders")

	if err := seeds.Run(); err != nil {
		logger.Log.Fatalf("error when runing seeders: %v", err)
	}

	logger.Log.Info("seeders executed correctly")
}
