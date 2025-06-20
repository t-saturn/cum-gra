package database

import (
	"fmt"

	"github.com/t-saturn/central-user-manager/internal/infrastructure/config"
	"github.com/t-saturn/central-user-manager/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	DB = db

	logger.Log.Info("Conexión exitosa a la base de datos PostgreSQL")
}
