package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Constantes para estados y tipos
const (
	// Métodos de autenticación
	AuthMethodCredentials = "credentials"
	AuthMethodToken       = "token"

	// Estados de autenticación
	AuthStatusPending = "pending"
	AuthStatusSuccess = "success"
	AuthStatusInvalid = "invalid_credentials"
	AuthStatusFailed  = "failed"
	AuthStatusExpired = "expired"

	// Estados de tokens
	TokenStatusActive  = "active"
	TokenStatusRevoked = "revoked"
	TokenStatusExpired = "expired"
	TokenStatusInvalid = "invalid"

	// Tipos de tokens
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	// Estados de sesiones
	SessionStatusActive   = "active"
	SessionStatusInactive = "inactive"
	SessionStatusRevoked  = "revoked"
	SessionStatusExpired  = "expired"

	// Tipos de logout
	LogoutTypeUserInitiated  = "user_initiated"
	LogoutTypeSessionExpired = "session_expired"
	LogoutTypeAdminRevoked   = "admin_revoked"
	LogoutTypeSecurityBreach = "security_breach"

	// Tipos de dispositivos
	DeviceTypeDesktop = "desktop"
	DeviceTypeMobile  = "mobile"
	DeviceTypeTablet  = "tablet"
	DeviceTypeOther   = "other"
)

// --- TIPOS BÁSICOS ---
// Coordenadas geográficas [longitud, latitud]
type Coordinates [2]float64

// Información de ubicación básica
type Location struct {
	Country     string      `bson:"country,omitempty"`
	City        string      `bson:"city,omitempty"`
	Coordinates Coordinates `bson:"coordinates,omitempty"`
}

// Información de ubicación detallada
type LocationDetail struct {
	Country      string      `bson:"country,omitempty"`
	CountryCode  string      `bson:"country_code,omitempty"`
	Region       string      `bson:"region,omitempty"`
	City         string      `bson:"city,omitempty"`
	Coordinates  Coordinates `bson:"coordinates,omitempty"`
	ISP          string      `bson:"isp,omitempty"`
	Organization string      `bson:"organization,omitempty"`
}

// Información básica del dispositivo
type DeviceInfo struct {
	UserAgent      string    `bson:"user_agent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"device_id,omitempty"`
	BrowserName    string    `bson:"browser_name,omitempty"`
	BrowserVersion string    `bson:"browser_version,omitempty"`
	OS             string    `bson:"os,omitempty"`
	OSVersion      string    `bson:"os_version,omitempty"`
	DeviceType     string    `bson:"device_type,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}

// Información extendida del dispositivo para sesiones
type SessionDeviceInfo struct {
	UserAgent      string          `bson:"user_agent,omitempty"`
	IP             string          `bson:"ip,omitempty"`
	DeviceID       string          `bson:"device_id,omitempty"`
	BrowserName    string          `bson:"browser_name,omitempty"`
	BrowserVersion string          `bson:"browser_version,omitempty"`
	OS             string          `bson:"os,omitempty"`
	OSVersion      string          `bson:"os_version,omitempty"`
	DeviceType     string          `bson:"device_type,omitempty"`
	Timezone       string          `bson:"timezone,omitempty"`
	Language       string          `bson:"language,omitempty"`
	Location       *LocationDetail `bson:"location,omitempty"`
}

// --- CAPTCHA ---
// Log de validación de CAPTCHA
type CaptchaLog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Token       string             `bson:"token"`
	Success     bool               `bson:"success"`
	ChallengeTS time.Time          `bson:"challenge_ts"`
	Hostname    string             `bson:"hostname"`
	Action      string             `bson:"action,omitempty"`
	CustomData  string             `bson:"cdata,omitempty"`
	ErrorCodes  []string           `bson:"error_codes,omitempty"`
	RemoteIP    string             `bson:"remote_ip"`
	CreatedAt   time.Time          `bson:"created_at"`

	// Relaciones
	UserID        primitive.ObjectID  `bson:"user_id,omitempty"`
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`
	AuthAttemptID *primitive.ObjectID `bson:"auth_attempt_id,omitempty"`
	ApplicationID string              `bson:"application_id,omitempty"`
}

// --- RESPUESTA DE VALIDACIÓN ---
// Respuesta del servicio de validación
type ValidationResponse struct {
	UserID          primitive.ObjectID `bson:"user_id,omitempty"`
	ServiceResponse string             `bson:"service_response,omitempty"`
	ValidatedBy     string             `bson:"validated_by,omitempty"`
	ValidationTime  int                `bson:"validation_time,omitempty"` // tiempo en ms
}

// --- MÉTRICAS ---
// Métricas de sesión
type SessionMetrics struct {
	TotalRequests        int    `bson:"total_requests"`
	TotalTokenRefreshes  int    `bson:"total_token_refreshes"`
	ApplicationsAccessed int    `bson:"applications_accessed"`
	LastApplicationUsed  string `bson:"last_application_used,omitempty"`
	SessionDuration      int64  `bson:"session_duration"` // en segundos
}
