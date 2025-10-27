package migrations

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"central-user-manager/internal/config"
	"central-user-manager/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func HandleMigration() {
	migrationsPath := flag.String("path", "internal/database/migrations", "Directorio de migraciones")
	cmd := flag.String("cmd", "up", "Comando: up, down, force, version, drop")
	steps := flag.Int("steps", 0, "Cantidad de pasos para up/down")
	version := flag.Int("version", 0, "Versión para comando force")
	flag.Parse()

	sqlDB, err := config.DB.DB()
	if err != nil {
		logger.Log.Fatal("Error obteniendo instancia sql.DB:", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		logger.Log.Fatal("Error creando el driver de migración:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+*migrationsPath, "postgres", driver)
	if err != nil {
		logger.Log.Fatal("Error inicializando instancia de migración:", err)
	}

	commands := map[string]func(*migrate.Migrate) error{
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
		"reset": func(_ *migrate.Migrate) error {
			var content []byte
			content, err = os.ReadFile("internal/database/clean/reset_data.sql")
			if err != nil {
				return fmt.Errorf("no se pudo leer el script reset_data.sql: %w", err)
			}

			sqlDB, err = config.DB.DB()
			if err != nil {
				return fmt.Errorf("error obteniendo instancia sql.DB: %w", err)
			}
			_, err = sqlDB.Exec(string(content))
			if err != nil {
				return fmt.Errorf("error ejecutando reset_data.sql: %w", err)
			}

			logger.Log.Info("Datos reseteados exitosamente.")
			return nil
		},
		"force": func(m *migrate.Migrate) error {
			if *version == 0 {
				return fmt.Errorf("debes indicar una versión con -version")
			}
			return m.Force(*version)
		},
		"version": func(m *migrate.Migrate) error {
			v, dirty, versionErr := m.Version()
			if versionErr != nil {
				return versionErr
			}
			logger.Log.Infof("Versión actual: %d (dirty: %t)", v, dirty)
			return nil
		},
		"drop": func(m *migrate.Migrate) error {
			return m.Drop()
		},
	}

	action, exists := commands[*cmd]
	if !exists {
		logger.Log.Fatal("Comando inválido. Usa: up, down, force, version, drop, reset")
	}

	err = action(m)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Log.Fatalf("Error ejecutando migración '%s': %v", *cmd, err)
	}

	logger.Log.Infof("Migración '%s' ejecutada correctamente", *cmd)
}
