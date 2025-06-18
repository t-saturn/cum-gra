package migrations

import (
	"log"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateUserApplicationRelations() {
	// Ahora que users, applications y application_roles existen,
	// podemos crear user_application_roles
	err := database.DB.AutoMigrate(&domain.UserApplicationRole{})
	if err != nil {
		log.Fatalf("Error al migrar la tabla de roles de usuario por aplicación: %v", err)
	}

	log.Println("Migración de tabla de relaciones usuario-aplicación (user_application_roles) completada")
}
