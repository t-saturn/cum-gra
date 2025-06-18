package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateApplicationsTables() {
	// Migrar applications primero
	err := database.DB.AutoMigrate(&domain.Application{})
	if err != nil {
		log.Fatalf("Error al migrar la tabla de aplicaciones: %v", err)
	}

	// Luego migrar application_roles
	err = database.DB.AutoMigrate(&domain.ApplicationRole{})
	if err != nil {
		log.Fatalf("Error al migrar la tabla de roles de aplicaciones: %v", err)
	}

	log.Println("Migración de tablas de aplicaciones (applications, application_roles) completada")
}
