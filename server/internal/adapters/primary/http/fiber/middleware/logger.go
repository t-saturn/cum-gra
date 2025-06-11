// File: internal/adapters/primary/http/fiber/middleware/logger.go
package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware() fiber.Handler {
	// Crea o abre el archivo log
	file, err := os.OpenFile("logs/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return logger.New(logger.Config{
		Output:     file,
		TimeFormat: time.RFC3339,
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
	})
}
