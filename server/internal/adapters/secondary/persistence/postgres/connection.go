package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Error al conectar GORM: %w", err)
	}

	log.Println("Conectado a la base de datos PostgreSQL con GORM")

	// Al final de InitPostgres
	DB.AutoMigrate(&models.UserModel{})
	return nil
}
