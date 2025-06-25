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
	UserAgent      string    `bson:"userAgent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"deviceId,omitempty"`
	BrowserName    string    `bson:"browserName,omitempty"`
	BrowserVersion string    `bson:"browserVersion,omitempty"`
	OS             string    `bson:"os,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}

// Respuesta de validación
type ValidationResponse struct {
	UserID          primitive.ObjectID `bson:"userId,omitempty"`
	ServiceResponse string             `bson:"serviceResponse,omitempty"`
	ValidatedBy     string             `bson:"validatedBy,omitempty"`
}

// Modelo principal
type UserCredential struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty"`
	Email              string              `bson:"email"`
	PasswordHash       string              `bson:"passwordHash"`
	Status             string              `bson:"status"` // pending, correct, invalid, expired
	ApplicationID      string              `bson:"applicationId"`
	ApplicationURL     string              `bson:"applicationUrl"`
	DeviceInfo         DeviceInfo          `bson:"deviceInfo"`
	CreatedAt          time.Time           `bson:"createdAt"`
	ExpiresAt          time.Time           `bson:"expiresAt"`
	ValidatedAt        *time.Time          `bson:"validatedAt,omitempty"`
	ValidationResponse *ValidationResponse `bson:"validationResponse,omitempty"`
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
