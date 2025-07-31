package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
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
	errResp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    code,
		Message: message,
	}
	return c.Status(status).JSON(errResp)
}

// JSONResponse envuelve automáticamente tu ResponseDTO[T].
// message y data pueden omitirse (data puede ser el valor cero de T).
func JSONResponse[T any](c fiber.Ctx, status int, success bool, message string, data T) error {
	resp := dto.ResponseDTO[T]{
		Success: success,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(resp)
}
