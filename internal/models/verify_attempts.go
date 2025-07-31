package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyAttempt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Email         string             `bson:"email"`
	UserID        string             `bson:"user_id,omitempty"`
	ApplicationID string             `bson:"application_id"`

	Status         string     `bson:"status"` // "success", "failed", "locked", etc.
	Reason         string     `bson:"reason"` // "invalid_password", "user_not_found", etc.
	Method         string     `bson:"method"` // "credentials"
	CreatedAt      time.Time  `bson:"created_at"`
	ValidatedAt    *time.Time `bson:"validated_at,omitempty"`
	ValidationTime int64      `bson:"validation_time"` // en ms

	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	ValidationResponse *ValidationResponse `bson:"validation_response,omitempty"`
}
