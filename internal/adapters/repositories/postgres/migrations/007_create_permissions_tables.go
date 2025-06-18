package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreatePermissionsTables() {
	err := database.DB.AutoMigrate(
		&domain.ModuleRolePermission{},
	)
	if err != nil {
		log.Fatalf("Error al migrar la tabla de permisos: %v", err)
	}

	log.Println("Migración de tabla de permisos (module_role_permissions) completada")
}