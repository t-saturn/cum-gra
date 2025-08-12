package dto

import "time"

// TokensDTO agrupa los tokens de acceso y refresh.
type TokensDTO struct {
	AccessToken  TokenDetailDTO `json:"access_token"`
	RefreshToken TokenDetailDTO `json:"refresh_token"`
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
