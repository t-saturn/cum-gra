package dto

import "time"

type AuthVerifyRequestDTO struct {
	Email         *string       `json:"email,omitempty" validate:"omitempty,email"`
	DNI           *string       `json:"dni,omitempty" validate:"omitempty,len=8,numeric"`
	Password      string        `json:"password" validate:"required"`
	ApplicationID string        `json:"application_id" validate:"required,uuid4"`
	DeviceInfo    DeviceInfoDTO `json:"device_info" validate:"required"`
	CaptchaToken  *string       `json:"captcha_token,omitempty"`
}

// AuthVerifyResponseDTO define el contenido de "data" para el endpoint /auth/verify.
type AuthVerifyResponseDTO struct {
	AttemptID          string                `json:"attempt_id"`
	UserID             string                `json:"user_id,omitempty"`
	Status             string                `json:"status"`
	ValidatedAt        time.Time             `json:"validated_at"`
	ValidationResponse ValidationResponseDTO `json:"validation_response"`
}

// AuthLoginRequestDTO define la estructura de la petición para /auth/login.
type AuthLoginRequestDTO struct {
	Email         string        `json:"email" validate:"required,email"`
	Password      string        `json:"password" validate:"required,min=6"`
	ApplicationID string        `json:"application_id" validate:"required"`
	RememberMe    bool          `json:"remember_me"`
	DeviceInfo    DeviceInfoDTO `json:"device_info" validate:"required"`
}

// AuthLoginResponseDTO define la parte "data" de la respuesta para /auth/login.
type AuthLoginResponseDTO struct {
	UserID    string     `json:"user_id"`
	Session   SessionDTO `json:"session"`
	Tokens    TokensDTO  `json:"tokens"`
	AttemptID string     `json:"attempt_id"`
}

// LogoutRequestDTO define la petición para /auth/logout
type LogoutRequestDTO struct {
	Token      string `json:"token" validate:"required"`
	SessionID  string `json:"session_id" validate:"required,uuid4"`
	LogoutType string `json:"logout_type" validate:"required,oneof=user_logout refresh_token session_expired admin_revoked"`
}

// LogoutResponseDTO define la parte "data" de la respuesta para /auth/logout
type LogoutResponseDTO struct {
	SessionID     string    `json:"session_id"`
	RevokedAt     time.Time `json:"revoked_at"`
	TokensRevoked []string  `json:"tokens_revoked"`
}
