package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/server/pkg/config"

	_ "github.com/lib/pq"
	fiberadapter "github.com/t-saturn/central-user-manager/server/internal/adapters/primary/http/fiber"
)

func main() {
	// Cargar .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se pudo cargar el archivo .env, se usará configuración por defecto o variables del sistema")
	}

	// Cargar configuración
	cfg := config.LoadConfig()

	// Conexión a la base de datos
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Error haciendo ping a la base de datos: %v", err)
	}

	fmt.Println("✅ Conectado a la base de datos PostgreSQL")

	fiberServer := fiberadapter.NewFiberServer()
log.Fatal(fiberServer.Listen(":" + cfg.Port))

	// Rutas base
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Iniciar servidor
	port := cfg.Port
	fmt.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	log.Fatal(app.Listen(":" + port))
}
