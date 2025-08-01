package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/middlewares"
	"github.com/t-saturn/auth-service-server/internal/routes"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/validator"
)

func main() {
	// Cargar variables de entorno
	_ = godotenv.Load()

	// Inicializar logger
	logger.InitLogger()
	logger.Log.Info("Iniciando servidor...")

	// Cargar configuración y conectar a bases de datos
	config.LoadConfig()
	config.ConnectPostgres()
	config.ConnectMongo()

	// Inicializar validador
	if err := validator.InitValidator(); err != nil {
		logger.Log.Fatalf("Error al inicializar el validador: %v", err)
	}

	// Crear aplicación Fiber
	app := fiber.New()

	// Configurar middlewares
	app.Use(middlewares.CORSMiddleware())
	app.Use(middlewares.LoggerMiddleware())

	// Registrar rutas (sin dependencias adicionales)
	routes.RegisterRoutes(app)

	// Iniciar servidor
	port := config.GetConfig().Server.ServerPort
	logger.Log.Infof("Servidor escuchando en http://localhost:%s", port)

	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	// Desconectar MongoDB al cerrar
	defer config.DisconnectMongo()
}
