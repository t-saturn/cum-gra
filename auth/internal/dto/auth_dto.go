package dto

import (
	"time"
)

// AuthLoginRequestDTO define la estructura de la petición para /auth/login.
type AuthLoginRequestDTO struct {
	Email         string        `json:"email" validate:"required,email"`
	Password      string        `json:"password" validate:"required,min=6"`
	ApplicationID string        `json:"application_id" validate:"required"`
	DeviceInfo    DeviceInfoDTO `json:"device_info" validate:"required"`
}

type AuthLoginResponseDTO struct {
	UserID    string         `json:"user_id"`
	Session   SessionDTO     `json:"session"`
	Tokens    TokensLoginDTO `json:"tokens"`
	AttemptID string         `json:"attempt_id"`
}

// LogoutRequestDTO define la petición para /auth/logout
type LogoutRequestDTO struct {
	LogoutType string `validate:"required,oneof=user_logout refresh_token session_expired admin_revoked forbidden_revoke"`
	SessionID  string `json:"session_id" validate:"omitempty,uuid4"`
}

// LogoutResponseDTO define la parte "data" de la respuesta para /auth/logout
type LogoutResponseDTO struct {
	SessionID     string    `json:"session_id"`
	RevokedAt     time.Time `json:"revoked_at"`
	TokensRevoked []string  `json:"tokens_revoked"`
}

type SessionDTO struct {
	SessionID string    `json:"session_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
