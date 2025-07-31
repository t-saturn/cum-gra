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
