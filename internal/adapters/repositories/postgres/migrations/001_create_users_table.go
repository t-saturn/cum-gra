package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateUsersTable() {
	err := database.DB.AutoMigrate(
		&domain.StructuralPosition{},
		&domain.OrganicUnit{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas iniciales: %v", err)
	}

	log.Println("Migración de tablas users, structural_positions y organic_units completada")
}
