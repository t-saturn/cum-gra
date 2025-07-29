package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

// NowUTC devuelve la hora actual en UTC.
func NowUTC() time.Time {
	return time.Now().UTC()
}

// ParseISOTime parsea una cadena en formato ISO 8601 (RFC3339) a time.Time.
func ParseISOTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// JSON serializa payload a JSON y escribe el status HTTP.
func JSON(c fiber.Ctx, status int, payload interface{}) error {
	return c.Status(status).JSON(payload)
}

// JSONError envía un error con un código de aplicación y un mensaje.
func JSONError(c fiber.Ctx, status int, code, message string) error {
	// Puedes definir aquí tu propia estructura, o usar un map si te basta:
	errResp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    code,
		Message: message,
	}
	return c.Status(status).JSON(errResp)
}
