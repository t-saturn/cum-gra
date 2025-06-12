package database

import (
	"fmt"

	"github.com/t-saturn/central-user-manager/server/pkg/config"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = db
	logger.Log.Info("Connection to the database established successfully")
}
