package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sesión de usuario
type Session struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SessionID string             `bson:"session_id"` // UUID único de la sesión
	UserID    string             `bson:"user_id"`

	// Estado de la sesión
	Status       string     `bson:"status"`    // active, inactive, revoked, expired
	IsActive     bool       `bson:"is_active"` // Estado rápido de consulta
	MaxRefreshAt time.Time  `bson:"max_refresh_at"`
	CreatedAt    time.Time  `bson:"created_at"`
	LastActivity time.Time  `bson:"last_activity"`
	ExpiresAt    time.Time  `bson:"expires_at"`
	RevokedAt    *time.Time `bson:"revoked_at,omitempty"`

	// Información del dispositivo y ubicación
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Relaciones
	AuthAttemptID   *primitive.ObjectID  `bson:"auth_attempt_id,omitempty"`   // Intento que creó la sesión
	ParentSessionID *primitive.ObjectID  `bson:"parent_session_id,omitempty"` // Sesión padre (si aplica)
	ActiveTokens    []primitive.ObjectID `bson:"active_tokens,omitempty"`     // Tokens activos de esta sesión

	// Información de revocación
	RevokedBy        *primitive.ObjectID `bson:"revoked_by,omitempty"`        // Usuario/admin que revocó
	RevocationReason string              `bson:"revocation_reason,omitempty"` // Razón de revocación
}
