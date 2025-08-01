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

// ListSessionsRequestDTO agrupa los query params de GET /auth/sessions.
type ListSessionsRequestDTO struct {
	Status *string `query:"status"` // "active", "inactive", "revoked", "expired"
	Limit  *int    `query:"limit"`  // máximo resultados (default 10)
	Offset *int    `query:"offset"` // paginación (default 0)
}

// SessionSummaryDTO describe cada sesión en la lista.
type SessionSummaryDTO struct {
	SessionID    string        `json:"session_id"`
	Status       string        `json:"status"`
	IsActive     bool          `json:"is_active"`
	CreatedAt    time.Time     `json:"created_at"`
	LastActivity time.Time     `json:"last_activity"`
	ExpiresAt    time.Time     `json:"expires_at"`
	DeviceInfo   DeviceInfoDTO `json:"device_info"`
	IsCurrent    bool          `json:"is_current"`
}

// ListSessionsResponseDTO envuelve la respuesta de GET /auth/sessions.
type ListSessionsResponseDTO struct {
	Sessions []SessionSummaryDTO `json:"sessions"`
	Total    int64               `json:"total"`
	Limit    int                 `json:"limit"`
	Offset   int                 `json:"offset"`
}

// SessionRevokeRequestDTO define la petición para DELETE /auth/sessions/{session_id}
type SessionRevokeRequestDTO struct {
	Reason          string `json:"reason" validate:"required"`
	RevokeAllTokens bool   `json:"revoke_all_tokens"`
}

// SessionRevokeResponseDTO define la parte "data" de la respuesta para DELETE /auth/sessions/{session_id}
type SessionRevokeResponseDTO struct {
	SessionID        string    `json:"session_id"`
	RevokedAt        time.Time `json:"revoked_at"`
	RevocationReason string    `json:"revocation_reason"`
	TokensRevoked    []string  `json:"tokens_revoked"`
}
