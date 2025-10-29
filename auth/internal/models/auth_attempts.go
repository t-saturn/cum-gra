package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Intento de autenticación
type AuthAttempt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Method        string             `bson:"method"` // "credentials" | "token"
	Status        string             `bson:"status"` // "pending", "success", "failed", "expired"
	ApplicationID string             `bson:"application_id"`

	// Campos para method = "credentials"
	Email string `bson:"email,omitempty"`

	// Campos para method = "token"
	Token string `bson:"token,omitempty"`

	// Información del dispositivo
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Timestamps
	CreatedAt   time.Time  `bson:"created_at"`
	ValidatedAt *time.Time `bson:"validated_at,omitempty"`
	ExpiresAt   *time.Time `bson:"expires_at,omitempty"`

	// Respuesta de validación
	ValidationResponse *ValidationResponse `bson:"validation_response,omitempty"`

	// Relaciones
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`      // Sesión asociada (si existe)
	CaptchaLogID  *primitive.ObjectID `bson:"captcha_log_id,omitempty"`  // CAPTCHA validado
	ParentTokenID *primitive.ObjectID `bson:"parent_token_id,omitempty"` // Token padre (para refresh)
}
