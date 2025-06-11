package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupLogger() fiber.Handler {
	// Crear carpeta si no existe
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			log.Fatalf("Failed to create logs directory: %v", err)
		}
	}

	// Crear archivo con fecha actual
	fileName := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile("logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("No se pudo abrir archivo de logs: %v", err)
	}

	// Middleware de logger de Fiber
	return logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     file, // logs a archivo
	})
}
