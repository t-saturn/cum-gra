// Package main ejecuta las migraciones de la base de datos utilizando el motor configurado.
package main

import (
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/config"
	"github.com/t-saturn/central-user-manager/internal/database/migrations"
	"github.com/t-saturn/central-user-manager/pkg/logger"
)

func main() {
	logger.InitLogger()

	_ = godotenv.Load()
	config.LoadConfig()
	config.ConnectDB()

	migrations.HandleMigration()
}
