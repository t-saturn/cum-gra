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
		logger.Log.Warn("No se pudo cargar el archivo .env, usando entorno actual")
	}

	logger.InitLogger()

	config.LoadConfig()
	database.Connect()

	logger.Log.Info("Ejecutando seeders...")

	if err := seeds.Run(); err != nil {
		logger.Log.Fatalf("Error al ejecutar seeders: %v", err)
	}

	logger.Log.Info("Seeders ejecutados correctamente")
}
