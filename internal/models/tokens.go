package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token unificado (access y refresh)
type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TokenID   string             `bson:"token_id"`      // UUID único del token
	TokenHash string             `bson:"token_hash"`    // hash (p.ej. sha256 en hex/base64url)
	Alg       string             `bson:"alg,omitempty"` // "RS256"

	UserID    string `bson:"user_id"`
	SessionID string `bson:"session_id,omitempty"`

	Status    string `bson:"status"`     // active, revoked, expired
	TokenType string `bson:"token_type"` // access, refresh

	IssuedAt  time.Time  `bson:"issued_at"`
	ExpiresAt time.Time  `bson:"expires_at"`
	LastUsed  *time.Time `bson:"last_used,omitempty"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	RevokedAt *time.Time `bson:"revoked_at,omitempty"`

	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	ParentTokenID *primitive.ObjectID  `bson:"parent_token_id,omitempty"`
	PairedTokenID *primitive.ObjectID  `bson:"paired_token_id,omitempty"`
	ChildTokens   []primitive.ObjectID `bson:"child_tokens,omitempty"`

	Reason       string `bson:"reason,omitempty"`
	RevokedBy    string `bson:"revoked_by,omitempty"`
	RevokedByApp string `bson:"revoked_by_app,omitempty"`
}
