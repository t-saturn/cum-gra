package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Información de geolocalización
type LocationDetail struct {
	Country      string      `bson:"country,omitempty"`
	CountryCode  string      `bson:"country_code,omitempty"`
	Region       string      `bson:"region,omitempty"`
	City         string      `bson:"city,omitempty"`
	Coordinates  Coordinates `bson:"coordinates,omitempty"`
	ISP          string      `bson:"isp,omitempty"`
	Organization string      `bson:"organization,omitempty"`
}

// Información extendida del dispositivo
type SessionDeviceInfo struct {
	UserAgent      string          `bson:"user_agent,omitempty"`
	IP             string          `bson:"ip,omitempty"`
	DeviceID       string          `bson:"device_id,omitempty"`
	BrowserName    string          `bson:"browser_name,omitempty"`
	BrowserVersion string          `bson:"browser_version,omitempty"`
	OS             string          `bson:"os,omitempty"`
	OSVersion      string          `bson:"os_version,omitempty"`
	DeviceType     string          `bson:"device_type,omitempty"` // desktop, mobile, tablet
	Timezone       string          `bson:"timezone,omitempty"`
	Language       string          `bson:"language,omitempty"`
	Location       *LocationDetail `bson:"location,omitempty"`
}

// Métricas de la sesión
type SessionMetrics struct {
	TotalRequests        int    `bson:"total_requests"`
	TotalTokenRefreshes  int    `bson:"total_token_refreshes"`
	ApplicationsAccessed int    `bson:"applications_accessed"`
	LastApplicationUsed  string `bson:"last_application_used,omitempty"`
	SessionDuration      int64  `bson:"session_duration"` // en segundos
}

// Modelo de sesión
type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	SessionID    string             `bson:"session_id"`
	UserID       primitive.ObjectID `bson:"user_id"`
	IsActive     bool               `bson:"is_active"`
	CreatedAt    time.Time          `bson:"created_at"`
	LastActivity time.Time          `bson:"last_activity"`
	ExpiresAt    time.Time          `bson:"expires_at"`

	DeviceInfo SessionDeviceInfo `bson:"device_info,omitempty"`
	Metrics    SessionMetrics    `bson:"metrics,omitempty"`
}

type CaptchaLog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Token       string             `bson:"token"`                 // Token del captcha resuelto
	Success     bool               `bson:"success"`               // ¿Validación exitosa?
	ChallengeTS time.Time          `bson:"challenge_ts"`          // Hora del desafío
	Hostname    string             `bson:"hostname"`              // Dominio desde el cual se resolvió
	Action      string             `bson:"action,omitempty"`      // Acción definida (captcha v3)
	CustomData  string             `bson:"cdata,omitempty"`       // Información contextual opcional
	ErrorCodes  []string           `bson:"error_codes,omitempty"` // Errores del proveedor CAPTCHA
	RemoteIP    string             `bson:"remote_ip"`             // IP del usuario
	CreatedAt   time.Time          `bson:"created_at"`            // Hora de creación

	// Relaciones
	UserID        primitive.ObjectID  `bson:"user_id,omitempty"`        // Usuario que resolvió el captcha
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`     // Sesión activa, si existe
	AuthLogID     *primitive.ObjectID `bson:"auth_log_id,omitempty"`    // Registro de log asociado
	ApplicationID string              `bson:"application_id,omitempty"` // App desde la que se resolvió
}
