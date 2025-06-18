package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/central-user-manager/internal/adapters/repositories/postgres/migrations"
	"github.com/central-user-manager/internal/infrastructure/config"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/central-user-manager/pkg/logger"
)

func main() {
	// Cargar archivo .env manualmente
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, se usará el entorno actual")
	}

	// Inicializar el logger antes de usarlo
	logger.InitLogger()

	// Cargar configuración y conectar a la BD
	config.LoadConfig()
	database.Connect()

	// Ejecutar migraciones en orden de dependencias
	migrations.CreateEnumsAndExtensions()
	migrations.CreateBaseEntitiesTables()
	migrations.CreateApplicationsTables()
	migrations.CreateModulesTables()
	migrations.CreatePermissionsTables()
	migrations.CreateRestrictionsTables()
	migrations.CreateHistoryTables()
	migrations.CreateIndexes()
}
