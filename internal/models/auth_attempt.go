package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Respuesta del validador
type ValidationResponse struct {
	UserID          primitive.ObjectID `bson:"userId,omitempty"`
	ServiceResponse string             `bson:"serviceResponse,omitempty"`
	ValidatedBy     string             `bson:"validatedBy,omitempty"`
}

// Modelo principal
type AuthAttempt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Method        string             `bson:"method"` // "credentials" | "token"
	Status        string             `bson:"status"` // "pending", "success", "failed", "expired"
	ApplicationID string             `bson:"applicationId"`

	// Si method = "credentials"
	Email        string `bson:"email,omitempty"`
	PasswordHash string `bson:"passwordHash,omitempty"`

	// Si method = "token"
	Token string `bson:"token,omitempty"` // idealmente guardado como hash

	DeviceInfo DeviceInfo `bson:"deviceInfo,omitempty"`

	CreatedAt   time.Time  `bson:"createdAt"`
	ValidatedAt *time.Time `bson:"validatedAt,omitempty"`

	ValidationResponse *ValidationResponse `bson:"validationResponse,omitempty"`
}
