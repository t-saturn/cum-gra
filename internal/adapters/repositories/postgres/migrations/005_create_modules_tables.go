package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateModulesTables() {
	err := database.DB.AutoMigrate(
		&domain.Module{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de módulos: %v", err)
	}

	log.Println("Migración de tabla de módulos (modules) completada")
}