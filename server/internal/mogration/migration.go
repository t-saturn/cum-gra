package main

import (
	"log"

	"github.com/t-saturn/central-user-manager/server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuración de la base de datos
	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ejecutar migraciones
	err = db.AutoMigrate(
		// Tablas básicas
		&models.StructuralPosition{},
		&models.OrganicUnit{},
		&models.User{},
		&models.Application{},

		// Sesiones y tokens
		&models.UserSession{},
		&models.OAuthToken{},

		// Roles y permisos
		&models.ApplicationRole{},
		&models.UserApplicationRole{},
		&models.Module{},
		&models.ModuleRolePermission{},
		&models.UserPermission{},

		// Seguridad
		&models.PasswordHistory{},
		&models.PasswordReset{},
		&models.TwoFactorSecret{},

		// Auditoría y configuración
		&models.AuditLog{},
		&models.ApplicationSetting{},
		&models.UserPreference{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}
