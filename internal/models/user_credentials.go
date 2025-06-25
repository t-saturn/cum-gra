package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Coordenadas geográficas
type Coordinates [2]float64

// Información de ubicación
type Location struct {
	Country     string      `bson:"country,omitempty"`
	City        string      `bson:"city,omitempty"`
	Coordinates Coordinates `bson:"coordinates,omitempty"`
}

// Información del dispositivo
type DeviceInfo struct {
	UserAgent      string    `bson:"user_agent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"device_id,omitempty"`
	BrowserName    string    `bson:"browser_name,omitempty"`
	BrowserVersion string    `bson:"browser_version,omitempty"`
	OS             string    `bson:"os,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}

// Respuesta de validación
type ValidationResponse struct {
	UserID          primitive.ObjectID `bson:"user_id,omitempty"`
	ServiceResponse string             `bson:"service_response,omitempty"`
	ValidatedBy     string             `bson:"validated_by,omitempty"`
}

// Modelo principal
type UserCredential struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty"`
	Email              string              `bson:"email"`
	PasswordHash       string              `bson:"password_hash"`
	Status             string              `bson:"status"` // pending, correct, invalid, expired
	ApplicationID      string              `bson:"application_id"`
	ApplicationURL     string              `bson:"application_url"`
	DeviceInfo         DeviceInfo          `bson:"device_info"`
	CreatedAt          time.Time           `bson:"created_at"`
	ExpiresAt          time.Time           `bson:"expires_at"`
	ValidatedAt        *time.Time          `bson:"validated_at,omitempty"`
	ValidationResponse *ValidationResponse `bson:"validation_response,omitempty"`
}

// NewUserCredential crea una instancia con tiempo de expiración a 5 minutos
func NewUserCredential(email, passwordHash, appID, appURL string, device DeviceInfo) *UserCredential {
	now := time.Now().UTC()
	expiration := now.Add(5 * time.Minute)

	return &UserCredential{
		ID:             primitive.NewObjectID(),
		Email:          email,
		PasswordHash:   passwordHash,
		Status:         "pending",
		ApplicationID:  appID,
		ApplicationURL: appURL,
		DeviceInfo:     device,
		CreatedAt:      now,
		ExpiresAt:      expiration,
	}
}
