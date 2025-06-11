package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber"
	"github.com/t-saturn/central-user-manager/server/internal/adapters/secondary/persistence/postgres"
)

func main() {
	if err := postgres.InitPostgres(); err != nil {
		log.Fatalf("Error al conectar la base de datos: %v", err)
	}

	fiber.StartFiberServer()
}
