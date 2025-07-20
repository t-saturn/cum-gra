package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ValidationResponse struct {
	UserID          primitive.ObjectID `bson:"user_id,omitempty"`
	ServiceResponse string             `bson:"service_response,omitempty"`
	ValidatedBy     string             `bson:"validated_by,omitempty"`
}

type AuthAttempt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Method        string             `bson:"method"` // "credentials" | "token"
	Status        string             `bson:"status"` // "pending", "success", "failed", "expired"
	ApplicationID string             `bson:"application_id"`

	// Si method = "credentials"
	Email        string `bson:"email,omitempty"`
	PasswordHash string `bson:"password_hash,omitempty"`

	// Si method = "token"
	Token string `bson:"token,omitempty"` // idealmente guardado como hash

	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	CreatedAt   time.Time  `bson:"created_at"`
	ValidatedAt *time.Time `bson:"validated_at,omitempty"`

	ValidationResponse *ValidationResponse `bson:"validation_response,omitempty"`
}
