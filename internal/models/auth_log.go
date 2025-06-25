package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Detalles específicos del evento
type AuthLogDetails struct {
	// Para credential_validation
	CredentialStatus string `bson:"credentialStatus,omitempty"` // pending, correct, invalid, expired
	ValidationTime   int    `bson:"validationTime,omitempty"`   // ms

	// Para token_validation
	TokenType    string `bson:"tokenType,omitempty"`
	Refreshed    bool   `bson:"refreshed,omitempty"`
	RefreshCount int    `bson:"refreshCount,omitempty"`

	// Para errores
	ErrorCode    string `bson:"errorCode,omitempty"`
	ErrorMessage string `bson:"errorMessage,omitempty"`

	// Para logout
	LogoutType string `bson:"logoutType,omitempty"` // user_initiated, session_expired...

	// Adicional
	UserAgent          string `bson:"userAgent,omitempty"`
	PreviousIP         string `bson:"previousIP,omitempty"`
	SuspiciousActivity bool   `bson:"suspiciousActivity,omitempty"`
}

// Modelo de log de autenticación
type AuthLog struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	UserID         primitive.ObjectID  `bson:"userId"`
	SessionID      *primitive.ObjectID `bson:"sessionId,omitempty"`
	CredentialID   string              `bson:"credentialId,omitempty"` // puede ser UUID
	TokenID        string              `bson:"tokenId,omitempty"`
	Action         string              `bson:"action"` // ver lista estándar
	Success        bool                `bson:"success"`
	ApplicationID  string              `bson:"applicationId"`
	ApplicationURL string              `bson:"applicationUrl,omitempty"`
	Details        AuthLogDetails      `bson:"details,omitempty"`
	DeviceInfo     DeviceInfo          `bson:"deviceInfo,omitempty"`
	Timestamp      time.Time           `bson:"timestamp"`
	ProcessingTime int                 `bson:"processingTime"` // en milisegundos
}

// credential_submit, credential_validation, credential_expired
// token_generate, token_validate, token_refresh, token_invalidate, token_expire
// session_create, session_extend, session_terminate, sso_login
// login_attempt, login_success, logout, forced_logout
// suspicious_activity, security_violation, device_change, location_change
