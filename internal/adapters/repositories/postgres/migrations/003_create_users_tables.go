package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateUsersTables() {
	err := database.DB.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de usuarios: %v", err)
	}

	log.Println("Migración de tabla de usuarios (users) completada")
}