package main

import (
	"log"

	"github.com/central-user-manager/internal/adapters/repositories/postgres/migrations"
	"github.com/central-user-manager/internal/infrastructure/config"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/central-user-manager/pkg/logger"
	"github.com/joho/godotenv"
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

	// Deshabilitar claves foráneas temporalmente
	migrations.DisableForeignKeyConstraints()

	// Ejecutar migraciones
	migrations.CreateEnumsAndExtensions()
	migrations.CreateIndependentTables()
	migrations.CreateUsersTables()
	migrations.CreateApplicationsTables()
	migrations.CreateModulesTables()
	migrations.CreatePermissionsTables()
	migrations.CreateRestrictionsTables()
	migrations.CreateHistoryTables()

	// Rehabilitar claves foráneas
	migrations.EnableForeignKeyConstraints()

	log.Println("Todas las migraciones completadas exitosamente")
}
