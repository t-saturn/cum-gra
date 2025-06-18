package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateApplicationsTables() {
	err := database.DB.AutoMigrate(
		&domain.Application{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de aplicaciones: %v", err)
	}

	// Luego migrar application_roles que depende de applications
	err = database.DB.AutoMigrate(
		&domain.ApplicationRole{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de roles de aplicaciones: %v", err)
	}

	// Finalmente migrar user_application_roles que depende de users, applications y application_roles
	err = database.DB.AutoMigrate(
		&domain.UserApplicationRole{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de roles de usuario por aplicación: %v", err)
	}

	log.Println("Migración de tablas de aplicaciones (applications, application_roles, user_application_roles) completada")
}