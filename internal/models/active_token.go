package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActiveToken struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	TokenID         string              `bson:"tokenId"`
	UserID          primitive.ObjectID  `bson:"userId"`
	SessionID       *primitive.ObjectID `bson:"sessionId,omitempty"`
	TokenHash       string              `bson:"tokenHash"`
	TokenType       string              `bson:"tokenType"` // access or refresh
	ApplicationID   string              `bson:"applicationId"`
	ApplicationURL  string              `bson:"application_url,omitempty"`
	DeviceInfo      map[string]string   `bson:"device_info,omitempty"` // simplificado
	IssuedAt        time.Time           `bson:"issued_at"`
	ExpiresAt       time.Time           `bson:"expires_at"`
	LastUsed        time.Time           `bson:"last_used,omitempty"`
	CreatedAt       time.Time           `bson:"created_at"`
	RefreshCount    int                 `bson:"refresh_count,omitempty"`
	MaxRefreshCount int                 `bson:"max_refresh_count,omitempty"`
}
