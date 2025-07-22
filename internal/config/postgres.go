package config

import (
	"fmt"
	"time"

	"github.com/t-saturn/auth-service-server/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

// ConnectPostgres establece la conexión a la base de datos PostgreSQL utilizando los datos de configuración.
func ConnectPostgres() {
	cfg := GetConfig() // Obtiene la configuración global

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("Error al conectar con PostgreSQL: %v", err)
	}

	// Configurar el pool de conexiones
	sqlDB, err := db.DB()
	if err != nil {
		logger.Log.Fatalf("Error al obtener DB de GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	PostgresDB = db
	logger.Log.Infof("Conexión exitosa establecida a PostgreSQL ")
}
