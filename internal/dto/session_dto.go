package dto

import "time"

// SessionDTO representa la sesión creada.
type SessionDTO struct {
	SessionID string    `json:"session_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// ListSessionsParams define filtros y paginación para listar sesiones.
type ListSessionsParamsDTO struct {
	// Estado de la sesión (por ejemplo, "active", "revoked").
	Status *string `json:"status,omitempty" query:"status"`

	// Indica si la sesión está activa.
	IsActive *bool `json:"isActive,omitempty" query:"isActive"`
}

// SessionMeRequestDTO representa la petición para GET /auth/session/me
// No lleva body; el token se toma de la cabecera Authorization.
type SessionMeRequestDTO struct{}

// SessionMeResponseDTO define la parte "data" de la respuesta para GET /auth/session/me.
type SessionMeResponseDTO struct {
	SessionID    string        `json:"session_id"`
	UserID       string        `json:"user_id"`
	Status       string        `json:"status"`
	IsActive     bool          `json:"is_active"`
	CreatedAt    time.Time     `json:"created_at"`
	LastActivity time.Time     `json:"last_activity"`
	ExpiresAt    time.Time     `json:"expires_at"`
	DeviceInfo   DeviceInfoDTO `json:"device_info"`
	ActiveTokens []string      `json:"active_tokens"`
}
