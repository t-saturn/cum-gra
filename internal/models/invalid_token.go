package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvalidToken struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	TokenID            string             `bson:"token_id"`
	TokenHash          string             `bson:"token_hash"`
	UserID             primitive.ObjectID `bson:"user_id"`
	ApplicationID      string             `bson:"application_id"`
	InvalidatedAt      time.Time          `bson:"invalidated_at"`
	InvalidationReason string             `bson:"invalidation_reason"` // user_logout, invalid_token, etc.
	InvalidatedBy      string             `bson:"invalidated_by"`      // user, admin, system
}
