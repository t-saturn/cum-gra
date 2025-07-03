package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/t-saturn/central-user-manager/internal/infrastructure/config"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
	"github.com/t-saturn/central-user-manager/pkg/logger"
)

type migrationAction func(m *migrate.Migrate) error

func main() {
	logger.InitLogger()

	if err := godotenv.Load(); err != nil {
		logger.Log.Println("No se pudo cargar el archivo .env, se usará el entorno actual")
	}

	config.LoadConfig()
	database.Connect()

	migrationsPath := flag.String("path", "internal/adapters/repositories/postgres/migrations", "Path to migrations directory")
	cmd := flag.String("cmd", "up", "Migration command: up, down, force, version, drop")
	steps := flag.Int("steps", 0, "Number of steps for up/down migration")
	version := flag.Int("version", 0, "Version to force migration to")
	flag.Parse()

	sqlDB, err := database.DB.DB()
	if err != nil {
		logger.Log.Fatal("Error getting underlying sql.DB:", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		logger.Log.Fatal("Error creating postgres driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+*migrationsPath, "postgres", driver)
	if err != nil {
		logger.Log.Fatal("Error creating migrate instance:", err)
	}

	// Map de comandos a funciones
	commands := map[string]migrationAction{
		"up": func(m *migrate.Migrate) error {
			if *steps > 0 {
				return m.Steps(*steps)
			}
			return m.Up()
		},
		"down": func(m *migrate.Migrate) error {
			if *steps > 0 {
				return m.Steps(-*steps)
			}
			return m.Down()
		},
		"force": func(m *migrate.Migrate) error {
			if *version == 0 {
				return fmt.Errorf("version is required for force command")
			}
			return m.Force(*version)
		},
		"version": func(m *migrate.Migrate) error {
			v, dirty, versionErr := m.Version()
			if versionErr != nil {
				return versionErr
			}
			logger.Log.Infof("Current migration version: %d, dirty: %t", v, dirty)
			return nil
		},
		"drop": func(m *migrate.Migrate) error {
			return m.Drop()
		},
	}

	action, exists := commands[*cmd]
	if !exists {
		logger.Log.Fatal("Comando inválido. Usa: up, down, force, version, drop")
	}

	err = action(m)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Log.Fatalf("Error ejecutando migración '%s': %v", *cmd, err)
	}

	logger.Log.Infof("Migración '%s' ejecutada correctamente", *cmd)
}
