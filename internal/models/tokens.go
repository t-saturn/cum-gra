package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token unificado (access y refresh)
type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TokenID   string             `bson:"token_id"`   // UUID único del token
	TokenHash string             `bson:"token_hash"` // Hash del token para verificación

	// Identificación
	UserID    string `bson:"user_id"`
	SessionID string `bson:"session_id,omitempty"`

	// Estado y tipo
	Status    string `bson:"status"`     // active, revoked, expired
	TokenType string `bson:"token_type"` // access, refresh

	// Timestamps
	IssuedAt  time.Time  `bson:"issued_at"`
	ExpiresAt time.Time  `bson:"expires_at"`
	LastUsed  *time.Time `bson:"last_used,omitempty"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	RevokedAt *time.Time `bson:"revoked_at,omitempty"`

	// Información del dispositivo
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Relaciones entre tokens
	ParentTokenID *primitive.ObjectID  `bson:"parent_token_id,omitempty"` // Token padre (para refresh)
	PairedTokenID *primitive.ObjectID  `bson:"paired_token_id,omitempty"` // Token access/refresh pareado
	ChildTokens   []primitive.ObjectID `bson:"child_tokens,omitempty"`    // Tokens hijos generados

	// Información de revocación
	Reason       string `bson:"reason,omitempty"`         // user_logout, invalid_token, security_breach, etc.
	RevokedBy    string `bson:"revoked_by,omitempty"`     // Usuario/admin que revocó
	RevokedByApp string `bson:"revoked_by_app,omitempty"` // Aplicación que revocó
}
