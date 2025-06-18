package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateRestrictionsTables() {
	err := database.DB.AutoMigrate(
		&domain.UserModuleRestriction{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de restricciones: %v", err)
	}

	log.Println("Migración de tabla de restricciones (user_module_restrictions) completada")
}