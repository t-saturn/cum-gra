package migrations

import (
	"log"

	"github.com/central-user-manager/internal/infrastructure/database"
)

// DisableForeignKeyConstraints desactiva temporalmente las claves foráneas
func DisableForeignKeyConstraints() {
	err := database.DB.Exec("SET session_replication_role = replica;").Error
	if err != nil {
		log.Printf("Advertencia: No se pudieron deshabilitar las claves foráneas: %v", err)
	} else {
		log.Println("Claves foráneas deshabilitadas temporalmente")
	}
}

// EnableForeignKeyConstraints reactiva las claves foráneas
func EnableForeignKeyConstraints() {
	err := database.DB.Exec("SET session_replication_role = DEFAULT;").Error
	if err != nil {
		log.Printf("Error: No se pudieron rehabilitar las claves foráneas: %v", err)
	} else {
		log.Println("Claves foráneas rehabilitadas")
	}
}
