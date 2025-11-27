package migrations

import (
	"flag"
	"fmt"

	"server/internal/config"
	"server/internal/models"
	"server/pkg/logger"

	"gorm.io/gorm"
)

func HandleMigration() {
	cmd := flag.String("cmd", "up", "Comando: up, down, reset")
	flag.Parse()

	db := config.DB
	if db == nil {
		logger.Log.Fatal("La conexión a la base de datos (config.DB) es nil")
	}

	var err error

	switch *cmd {
	case "up":
		err = migrateUp(db)
	case "down":
		err = migrateDown(db)
	case "reset":
		err = migrateReset(db)
	default:
		logger.Log.Fatalf("Comando inválido. Usa: up, down, reset")
	}

	if err != nil {
		logger.Log.Fatalf("Error ejecutando migración '%s': %v", *cmd, err)
	}

	logger.Log.Infof("Migración '%s' ejecutada correctamente", *cmd)
}

func migrateUp(db *gorm.DB) error {
	logger.Log.Info("Ejecutando AutoMigrate con GORM...")

	return db.AutoMigrate(
		&models.User{},
		&models.Application{},
		&models.ApplicationRole{},
		&models.Module{},
		&models.OrganicUnit{},
		&models.StructuralPosition{},
		&models.Ubigeo{},
		&models.UserDetail{},
		&models.UserApplicationRole{},
		&models.ModuleRolePermission{},
		&models.UserModuleRestriction{},
	)
}

func migrateDown(db *gorm.DB) error {
	logger.Log.Info("Eliminando todas las tablas del sistema...")

	// Primero hijos, luego padres
	tables := []interface{}{
		&models.UserModuleRestriction{},
		&models.ModuleRolePermission{},
		&models.UserApplicationRole{},
		&models.UserDetail{},
		&models.Module{},
		&models.ApplicationRole{},
		&models.Application{},
		&models.Ubigeo{},
		&models.StructuralPosition{},
		&models.OrganicUnit{},
		&models.User{},
	}

	for _, t := range tables {
		if err := db.Migrator().DropTable(t); err != nil {
			return fmt.Errorf("error al eliminar tabla %T: %w", t, err)
		}
	}

	return nil
}

func migrateReset(db *gorm.DB) error {
	logger.Log.Info("Reseteando datos (TRUNCATE) sin eliminar tablas...")

	query := `
TRUNCATE TABLE
    user_module_restrictions,
    module_role_permissions,
    user_application_roles,
    user_details,
    modules,
    application_roles,
    applications,
    ubigeos,
    structural_positions,
    organic_units,
    users
RESTART IDENTITY CASCADE;
`
	if err := db.Exec(query).Error; err != nil {
		return fmt.Errorf("error ejecutando TRUNCATE: %w", err)
	}

	return nil
}
