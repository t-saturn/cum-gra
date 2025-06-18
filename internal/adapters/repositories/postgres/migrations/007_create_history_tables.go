package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateHistoryTables() {
	err := database.DB.AutoMigrate(
		&domain.PasswordHistory{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de historial: %v", err)
	}

	log.Println("Migración de tabla de historial (password_history) completada")
}