package config

import (
	"fmt"

	"github.com/t-saturn/auth-service-server/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB representa la instancia global de conexión a PostgreSQL.
var DB *gorm.DB

// ConnectDB establece la conexión a la base de datos PostgreSQL utilizando los datos de configuración.
func ConnectDB() {
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

	DB = db
	logger.Log.Infof("Conectado a PostgreSQL en %s:%s", cfg.Postgres.Host, cfg.Postgres.Port)
}
