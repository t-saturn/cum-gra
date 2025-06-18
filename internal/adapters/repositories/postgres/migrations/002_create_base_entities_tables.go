package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateBaseEntitiesTables() {
	err := database.DB.AutoMigrate(
		&domain.StructuralPosition{},
		&domain.OrganicUnit{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas base: %v", err)
	}

	log.Println("Migración de tablas base (users, structural_positions, organic_units) completada")
}