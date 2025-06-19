// main.go
package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/central-user-manager/internal/infrastructure/config"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/central-user-manager/pkg/logger"
)

func main() {
	// Cargar archivo .env manualmente
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, se usará el entorno actual")
	}

	// Inicializar el logger antes de usarlo
	logger.InitLogger()

	// Cargar configuración y conectar a la BD
	config.LoadConfig()
	database.Connect()

	// Flags para comandos
	var (
		migrationsPath = flag.String("path", "internal/adapters/repositories/postgres/migrations", "Path to migrations directory")
		command        = flag.String("cmd", "up", "Migration command: up, down, force, version, drop")
		steps          = flag.Int("steps", 0, "Number of steps for up/down migration")
		version        = flag.Int("version", 0, "Version to force migration to")
	)
	flag.Parse()

	// Obtener conexión a la base de datos
	db := database.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting underlying sql.DB:", err)
	}

	// Configurar driver de migración
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating postgres driver:", err)
	}

	// Crear instancia de migrate
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+*migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Error creating migrate instance:", err)
	}

	// Ejecutar comando de migración
	switch *command {
	case "up":
		if *steps > 0 {
			err = m.Steps(*steps)
		} else {
			err = m.Up()
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error running up migration:", err)
		}
		log.Println("Migration up completed successfully")

	case "down":
		if *steps > 0 {
			err = m.Steps(-*steps)
		} else {
			err = m.Down()
		}
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error running down migration:", err)
		}
		log.Println("Migration down completed successfully")

	case "force":
		if *version == 0 {
			log.Fatal("Version is required for force command")
		}
		err = m.Force(*version)
		if err != nil {
			log.Fatal("Error forcing migration:", err)
		}
		log.Printf("Migration forced to version %d successfully", *version)

	case "version":
		version, dirty, e := m.Version()
		if err != nil {
			log.Fatal("Error getting migration version:", e)
		}
		log.Printf("Current migration version: %d, dirty: %t", version, dirty)

	case "drop":
		err = m.Drop()
		if err != nil {
			log.Fatal("Error dropping database:", err)
		}
		log.Println("Database dropped successfully")

	default:
		log.Fatal("Invalid command. Use: up, down, force, version, drop")
	}
}
