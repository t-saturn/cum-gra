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

	// Cargar configuración y establecer conexiones a bases de datos
	config.LoadConfig()
	config.ConnectPostgres()
	config.ConnectMongo()

	// Crear instancia de Fiber
	app := fiber.New()

	// Configurar middlewares
	app.Use(middlewares.CORSMiddleware())
	app.Use(middlewares.LoggerMiddleware())

	// Registrar rutas
	routes.RegisterRoutes(app)

	// Iniciar servidor
	port := config.GetConfig().Server.ServerPort
	logger.Log.Infof("Servidor escuchando en http://localhost:%s", port)

	// Manejar cierre graceful
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	// Cerrar conexión de MongoDB al terminar
	defer config.DisconnectMongo()
}
