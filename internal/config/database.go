// Package config contiene la configuración de entorno y conexión a base de datos.
package config

import (
	"fmt"

	"github.com/t-saturn/central-user-manager/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB representa la instancia global de conexión a la base de datos GORM.
var DB *gorm.DB

// ConnectDB establece la conexión a la base de datos PostgreSQL utilizando los datos de configuración.
func ConnectDB() {
	cfg := GetConfig()

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
