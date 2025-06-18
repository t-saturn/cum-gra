package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateIndependentTables() {
	err := database.DB.AutoMigrate(
		&domain.StructuralPosition{},
		&domain.OrganicUnit{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas independientes: %v", err)
	}

	log.Println("Migración de tablas independientes (structural_positions, organic_units) completada")
}