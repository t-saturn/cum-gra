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

// TokenValidationRequest representa el JSON recibido con el token.
type TokenValidationRequestDTO struct {
	Token         string `json:"token" validate:"required"`
	ApplicationID string `json:"application_id" validate:"required"`
}

type TokenValidationResponseDTO struct {
	UserID      string                            `json:"user_id"`
	TokenID     string                            `json:"token_id"`
	SessionID   string                            `json:"session_id"`
	Status      string                            `json:"status"`
	TokenType   string                            `json:"token_type"`
	TokenDetail TokenValidationDetailsResponseDTO `json:"token_detail"`
}

type TokenValidationDetailsResponseDTO struct {
	Valid     bool   `json:"valid"`
	Message   string `json:"message,omitempty"`
	Subject   string `json:"subject,omitempty"`
	IssuedAt  string `json:"issued_at,omitempty"`  // formato ISO8601
	ExpiresAt string `json:"expires_at,omitempty"` // formato ISO8601
	ExpiresIn int64  `json:"expires_in,omitempty"`
}
