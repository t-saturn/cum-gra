package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Modelo de token unificado
type Token struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	TokenID   string              `bson:"tokenId"` // jti
	TokenHash string              `bson:"tokenHash"`
	UserID    primitive.ObjectID  `bson:"userId"`
	SessionID *primitive.ObjectID `bson:"sessionId,omitempty"`

	Status    string    `bson:"status"`           // active, invalid, expired, revoked
	Reason    string    `bson:"reason,omitempty"` // user_logout, invalid_token, etc.
	UpdatedAt time.Time `bson:"updatedAt"`        // último cambio de estado

	TokenType string     `bson:"tokenType"` // access, refresh
	IssuedAt  time.Time  `bson:"issuedAt"`
	ExpiresAt time.Time  `bson:"expiresAt"`
	LastUsed  *time.Time `bson:"lastUsed,omitempty"`
	CreatedAt time.Time  `bson:"createdAt"`

	ApplicationID  string     `bson:"applicationId"`
	ApplicationURL string     `bson:"applicationUrl,omitempty"`
	DeviceInfo     DeviceInfo `bson:"deviceInfo,omitempty"`

	RefreshCount    int `bson:"refreshCount,omitempty"`
	MaxRefreshCount int `bson:"maxRefreshCount,omitempty"`
}
