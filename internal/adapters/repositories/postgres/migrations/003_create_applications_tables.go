package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateApplicationsTables() {
	err := database.DB.AutoMigrate(
		&domain.Application{},
		&domain.ApplicationRole{},
		&domain.UserApplicationRole{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas de aplicaciones: %v", err)
	}

	log.Println("Migración de tablas de aplicaciones (applications, application_roles, user_application_roles) completada")
}