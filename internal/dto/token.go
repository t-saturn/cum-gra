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

// TokenValidationRequestDTO representa el JSON recibido para validar un token.
type TokenValidationRequestDTO struct {
	Token     string `json:"token" validate:"required"`      // Token crudo (JWS completo)
	SessionID string `json:"session_id" validate:"required"` // ID de la sesión asociada
}

// TokenValidationResponseDTO es la respuesta al validar un token.
type TokenValidationResponseDTO struct {
	UserID      string                            `json:"user_id"`
	TokenID     string                            `json:"token_id"`
	SessionID   string                            `json:"session_id"`
	Status      string                            `json:"status"`
	TokenType   string                            `json:"token_type"`
	TokenDetail TokenValidationDetailsResponseDTO `json:"token_detail"`
}

// TokenValidationDetailsResponseDTO contiene información detallada del token.
type TokenValidationDetailsResponseDTO struct {
	Valid     bool   `json:"valid"`
	Message   string `json:"message"`
	Subject   string `json:"subject,omitempty"`
	IssuedAt  string `json:"issued_at,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
	ExpiresIn int64  `json:"expires_in,omitempty"` // segundos restantes
}

// AuthRefreshRequestDTO define la estructura de la petición para /auth/token/refresh
type AuthRefreshRequestDTO struct {
	RefreshToken string        `json:"token" validate:"required"` // JWS crudo
	SessionID    string        `json:"session_id" validate:"required"`
	DeviceInfo   DeviceInfoDTO `json:"device_info"` // opcional, si quieres actualizar metadata
}

// AuthRefreshResponseDTO define la parte "data" de la respuesta para /auth/token/refresh
type AuthRefreshResponseDTO struct {
	AccessToken  TokenDetailDTO `json:"access_token"`
	RefreshToken TokenDetailDTO `json:"refresh_token"`
	SessionID    string         `json:"session_id"`
	RefreshCount int            `json:"refresh_count"`
}
