package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
	"github.com/t-saturn/central-user-manager/server/pkg/logger"
)

func main() {
	// Cargar variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Logger
	logger.InitLogger()
	logger.Log.Info("Iniciando servidor HTTP...")

	// Base de datos
	database.InitDatabase()

	// Iniciar Fiber
	app := fiber.New()

	// Rutas simples por ahora
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Puerto desde ENV o por defecto 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log.Infof("Servidor escuchando en puerto %s", port)
	log.Fatal(app.Listen(":" + port))
}
