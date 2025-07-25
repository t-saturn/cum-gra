package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Modelo de token unificado
type Token struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	TokenID   string              `bson:"token_id"`
	TokenHash string              `bson:"token_hash"`
	UserID    primitive.ObjectID  `bson:"user_id"`
	SessionID *primitive.ObjectID `bson:"session_id,omitempty"`

	Status    string    `bson:"status"`           // active, invalid, expired, revoked
	Reason    string    `bson:"reason,omitempty"` // user_logout, invalid_token, etc.
	UpdatedAt time.Time `bson:"updated_at"`

	TokenType string     `bson:"token_type"` // access, refresh
	IssuedAt  time.Time  `bson:"issued_at"`
	ExpiresAt time.Time  `bson:"expires_at"`
	LastUsed  *time.Time `bson:"last_used,omitempty"`
	CreatedAt time.Time  `bson:"created_at"`

	ApplicationID  string     `bson:"application_id"`
	ApplicationURL string     `bson:"application_url,omitempty"`
	DeviceInfo     DeviceInfo `bson:"device_info,omitempty"`

	RefreshCount    int `bson:"refresh_count,omitempty"`
	MaxRefreshCount int `bson:"max_refresh_count,omitempty"`
}
