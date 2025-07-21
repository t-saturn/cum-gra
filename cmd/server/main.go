package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/middlewares"
	"github.com/t-saturn/auth-service-server/internal/routes"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

func main() {
	_ = godotenv.Load()

	logger.InitLogger()
	logger.Log.Info("Iniciando servidor...")

	config.LoadConfig()
	config.ConnectMongo()

	app := fiber.New()

	app.Use(middlewares.CORSMiddleware())
	app.Use(middlewares.LoggerMiddleware())

	routes.RegisterRoutes(app)

	// Ruta de prueba para verificar conexión a Mongo
	app.Get("/ping", func(c fiber.Ctx) error {
		if config.MongoClient == nil || config.MongoDatabase == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "MongoDB no conectado"})
		}
		return c.JSON(fiber.Map{"message": "Conectado a MongoDB correctamente"})
	})

	port := config.GetConfig().PORT
	logger.Log.Infof("Servidor escuchando en http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
