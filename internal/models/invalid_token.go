package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvalidToken struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	TokenID            string             `bson:"tokenId"`
	TokenHash          string             `bson:"tokenHash"`
	UserID             primitive.ObjectID `bson:"userId"`
	ApplicationID      string             `bson:"applicationId"`
	InvalidatedAt      time.Time          `bson:"invalidatedAt"`
	InvalidationReason string             `bson:"invalidationReason"` // user_logout, invalid_token, etc.
	InvalidatedBy      string             `bson:"invalidatedBy"`      // user, admin, system
}
