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
	CredentialID   string              `bson:"credential_id,omitempty"` // puede ser UUID
	TokenID        string              `bson:"token_id,omitempty"`
	Action         string              `bson:"action,omitempty"` // ver lista estándar
	Success        bool                `bson:"success,omitempty"`
	ApplicationID  string              `bson:"application_id,omitempty"`
	ApplicationURL string              `bson:"application_url,omitempty"`
	Details        AuthLogDetails      `bson:"details,omitempty"`
	DeviceInfo     DeviceInfo          `bson:"device_info,omitempty"`
	Timestamp      time.Time           `bson:"timestamp"`
	ProcessingTime int                 `bson:"processing_time,omitempty"` // en milisegundos
}

// credential_submit, credential_validation, credential_expired
// token_generate, token_validate, token_refresh, token_invalidate, token_expire
// session_create, session_extend, session_terminate, sso_login
// login_attempt, login_success, logout, forced_logout
// suspicious_activity, security_violation, device_change, location_change
