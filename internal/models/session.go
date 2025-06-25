package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Información de geolocalización
type LocationDetail struct {
	Country      string      `bson:"country,omitempty"`
	CountryCode  string      `bson:"countryCode,omitempty"`
	Region       string      `bson:"region,omitempty"`
	City         string      `bson:"city,omitempty"`
	Coordinates  Coordinates `bson:"coordinates,omitempty"`
	ISP          string      `bson:"isp,omitempty"`
	Organization string      `bson:"organization,omitempty"`
}

// Información extendida del dispositivo
type SessionDeviceInfo struct {
	UserAgent        string          `bson:"userAgent,omitempty"`
	IP               string          `bson:"ip,omitempty"`
	DeviceID         string          `bson:"deviceId,omitempty"`
	BrowserName      string          `bson:"browserName,omitempty"`
	BrowserVersion   string          `bson:"browserVersion,omitempty"`
	OS               string          `bson:"os,omitempty"`
	OSVersion        string          `bson:"osVersion,omitempty"`
	DeviceType       string          `bson:"deviceType,omitempty"` // desktop, mobile, tablet
	ScreenResolution string          `bson:"screenResolution,omitempty"`
	Timezone         string          `bson:"timezone,omitempty"`
	Language         string          `bson:"language,omitempty"`
	Location         *LocationDetail `bson:"location,omitempty"`
}

// Métricas de la sesión
type SessionMetrics struct {
	TotalRequests        int    `bson:"totalRequests"`
	TotalTokenRefreshes  int    `bson:"totalTokenRefreshes"`
	ApplicationsAccessed int    `bson:"applicationsAccessed"`
	LastApplicationUsed  string `bson:"lastApplicationUsed,omitempty"`
	SessionDuration      int64  `bson:"sessionDuration"` // en segundos
}

// Modelo de sesión
type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	SessionID    string             `bson:"sessionId"`
	UserID       primitive.ObjectID `bson:"userId"`
	IsActive     bool               `bson:"isActive"`
	CreatedAt    time.Time          `bson:"createdAt"`
	LastActivity time.Time          `bson:"lastActivity"`
	ExpiresAt    time.Time          `bson:"expiresAt"`

	DeviceInfo SessionDeviceInfo `bson:"deviceInfo,omitempty"`
	Metrics    SessionMetrics    `bson:"metrics,omitempty"`
}
