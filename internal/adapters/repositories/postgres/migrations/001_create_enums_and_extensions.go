package migrations

import (
	"log"

	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateEnumsAndExtensions() {
	// Crear extensión UUID
	err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Fatalf("Error al crear extensión uuid-ossp: %v", err)
	}

	// Crear ENUMs
	enums := []string{
		`DO $$ BEGIN
			CREATE TYPE status_enum AS ENUM ('active', 'suspended', 'deleted');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
		
		`DO $$ BEGIN
			CREATE TYPE application_status_enum AS ENUM ('active', 'suspended');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
		
		`DO $$ BEGIN
			CREATE TYPE permission_type_enum AS ENUM ('denied', 'read', 'write', 'admin');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
		
		`DO $$ BEGIN
			CREATE TYPE module_status_enum AS ENUM ('active', 'inactive');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
		
		`DO $$ BEGIN
			CREATE TYPE restriction_type_enum AS ENUM ('block_access', 'limit_permission');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
		
		`DO $$ BEGIN
			CREATE TYPE permission_level_enum AS ENUM ('denied', 'read', 'write', 'admin');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;`,
	}

	for _, enum := range enums {
		err := database.DB.Exec(enum).Error
		if err != nil {
			log.Fatalf("Error al crear enum: %v", err)
		}
	}

	log.Println("Migración de extensiones y enums completada")
}