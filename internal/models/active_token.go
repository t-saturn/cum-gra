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
	ApplicationURL  string              `bson:"applicationUrl,omitempty"`
	DeviceInfo      map[string]string   `bson:"deviceInfo,omitempty"` // simplificado
	IssuedAt        time.Time           `bson:"issuedAt"`
	ExpiresAt       time.Time           `bson:"expiresAt"`
	LastUsed        time.Time           `bson:"lastUsed,omitempty"`
	CreatedAt       time.Time           `bson:"createdAt"`
	RefreshCount    int                 `bson:"refreshCount,omitempty"`
	MaxRefreshCount int                 `bson:"maxRefreshCount,omitempty"`
}
