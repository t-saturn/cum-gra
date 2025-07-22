// Package dto define estructuras de datos utilizadas para la entrada y salida en las operaciones de la API.
package dto

// MessageResponse representa una respuesta simple con un mensaje.
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse representa una respuesta con un mensaje de error.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ValidationErrorResponse representa una respuesta con errores de validación por campo.
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// HealthResponse representa la respuesta del endpoint /health
type HealthResponse struct {
	Status  string `json:"status"`  // e.g. "ok"
	Message string `json:"message"` // texto adicional opcional
}
