package dto

import "time"

const (
	TokenStatusActive  = "active"
	TokenStatusExpired = "expired"
	TokenStatusRevoked = "revoked"
	TokenStatusUnknown = "unknown"

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

// IntrospectResponseDTO es la respuesta agregada de sesión + tokens.
type IntrospectResponseDTO struct {
	UserID           string    `json:"user_id"`
	SessionID        string    `json:"session_id"`
	Status           string    `json:"status"`             // estado de la sesión (active, revoked, expired, etc.)
	SessionExpiresAt string    `json:"session_expires_at"` // ISO8601
	Tokens           TokensDTO `json:"tokens"`
}

// TokensDTO agrupa los tokens de acceso y refresh.
type TokensDTO struct {
	AccessToken  TokenViewDTO `json:"access_token"`
	RefreshToken TokenViewDTO `json:"refresh_token"`
}

// TokenViewDTO representa el estado de un token específico.
type TokenViewDTO struct {
	TokenID     string                       `json:"token_id"`
	Status      string                       `json:"status"`       // active | expired | revoked | unknown
	TokenType   string                       `json:"token_type"`   // access | refresh
	TokenDetail IntrospectDetailsResponseDTO `json:"token_detail"` // detalles (claims/tiempos)
}

// TokenDetailDTO detalla un token individual.
type TokenDetailDTO struct {
	TokenID   string    `json:"token_id"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	ExpiresAt time.Time `json:"expires_at"`
}

type IntrospectQueryDTO struct {
	SessionID string `query:"session_id" json:"session_id" validate:"required"` // ID de la sesión asociada
}

type TokenIntrospectResponseDTO struct {
	UserID      string                       `json:"user_id"`
	TokenID     string                       `json:"token_id"`
	SessionID   string                       `json:"session_id"`
	Status      string                       `json:"status"`
	TokenType   string                       `json:"token_type"`
	TokenDetail IntrospectDetailsResponseDTO `json:"token_detail"`
}

// IntrospectDetailsResponseDTO contiene información detallada del token.
type IntrospectDetailsResponseDTO struct {
	Valid     bool   `json:"valid"`
	Message   string `json:"message"`
	Subject   string `json:"subject,omitempty"`
	IssuedAt  string `json:"issued_at,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
	ExpiresIn int64  `json:"expires_in,omitempty"` // segundos restantes
}

type AuthRefreshQueryDTO struct {
	SessionID string `query:"session_id" json:"session_id" validate:"required"`
}

type AuthRefreshResquestDTO struct {
	DeviceInfo DeviceInfoDTO `json:"device_info"`
}

type AuthRefreshResponseDTO struct {
	AccessToken  TokenDetailDTO `json:"access_token"`
	RefreshToken TokenDetailDTO `json:"refresh_token"`
	SessionID    string         `json:"session_id"`
	RefreshCount int            `json:"refresh_count"`
}

// TokenLoginDTO representa el token dentro de la respuesta de Login
type TokenLoginDTO struct {
	TokenID   string    `json:"token_id"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	ExpiresAt time.Time `json:"expires_at"`
}

type TokensLoginDTO struct {
	AccessToken  TokenLoginDTO `json:"access_token"`
	RefreshToken TokenLoginDTO `json:"refresh_token"`
}
