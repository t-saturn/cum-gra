package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Detalles específicos del evento
type AuthLogDetails struct {
	// Para credential_validation
	CredentialStatus string `bson:"credential_status,omitempty"` // pending, correct, invalid, expired
	ValidationTime   int    `bson:"validationTime,omitempty"`    // ms

	// Para token_validation
	TokenType    string `bson:"token_type,omitempty"`
	Refreshed    bool   `bson:"refreshed,omitempty"`
	RefreshCount int    `bson:"refresh_count,omitempty"`

	// Para errores
	ErrorCode    string `bson:"error_code,omitempty"`
	ErrorMessage string `bson:"error_message,omitempty"`

	// Para logout
	LogoutType string `bson:"logout_type,omitempty"` // user_initiated, session_expired...

	// Adicional
	UserAgent          string `bson:"user_agent,omitempty"`
	PreviousIP         string `bson:"previous_ip,omitempty"`
	SuspiciousActivity bool   `bson:"suspicious_activity,omitempty"`
}

// Modelo de log de autenticación
type AuthLog struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	UserID         primitive.ObjectID  `bson:"user_id,omitempty"`
	SessionID      *primitive.ObjectID `bson:"session_id,omitempty"`
	CredentialID   string              `bson:"credential_id,omitempty"`
	TokenID        string              `bson:"token_id,omitempty"`
	Action         string              `bson:"action,omitempty"` // login, token_validation, captcha, etc.
	Success        bool                `bson:"success,omitempty"`
	ApplicationID  string              `bson:"application_id,omitempty"`
	ApplicationURL string              `bson:"application_url,omitempty"`
	Timestamp      time.Time           `bson:"timestamp"`
	ProcessingTime int                 `bson:"processing_time,omitempty"`

	DeviceInfo DeviceInfo     `bson:"device_info,omitempty"`
	Details    AuthLogDetails `bson:"details,omitempty"`

	// Nuevo: relación con el resultado del captcha
	CaptchaLogID *primitive.ObjectID `bson:"captcha_log_id,omitempty"`
}
