package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token unificado (access y refresh)
type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TokenID   string             `bson:"token_id"`   // UUID único del token
	TokenHash string             `bson:"token_hash"` // Hash del token para verificación

	// Identificación
	UserID    primitive.ObjectID  `bson:"user_id"`
	SessionID *primitive.ObjectID `bson:"session_id,omitempty"`

	// Estado y tipo
	Status    string `bson:"status"`     // active, invalid, expired, revoked
	TokenType string `bson:"token_type"` // access, refresh

	// Timestamps
	IssuedAt  time.Time  `bson:"issued_at"`
	ExpiresAt time.Time  `bson:"expires_at"`
	LastUsed  *time.Time `bson:"last_used,omitempty"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	RevokedAt *time.Time `bson:"revoked_at,omitempty"`

	// Información de aplicación
	ApplicationID  string `bson:"application_id"`
	ApplicationURL string `bson:"application_url,omitempty"`

	// Información del dispositivo
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Para tokens refresh
	RefreshCount    int `bson:"refresh_count,omitempty"`     // Cuántas veces se ha usado para refresh
	MaxRefreshCount int `bson:"max_refresh_count,omitempty"` // Máximo permitido

	// Relaciones entre tokens
	ParentTokenID *primitive.ObjectID  `bson:"parent_token_id,omitempty"` // Token padre (para refresh)
	PairedTokenID *primitive.ObjectID  `bson:"paired_token_id,omitempty"` // Token access/refresh pareado
	ChildTokens   []primitive.ObjectID `bson:"child_tokens,omitempty"`    // Tokens hijos generados

	// Información de revocación
	Reason       string              `bson:"reason,omitempty"`         // user_logout, invalid_token, security_breach, etc.
	RevokedBy    *primitive.ObjectID `bson:"revoked_by,omitempty"`     // Usuario/admin que revocó
	RevokedByApp string              `bson:"revoked_by_app,omitempty"` // Aplicación que revocó

	// Información de seguridad
	RiskScore       int      `bson:"risk_score,omitempty"`       // Puntuación de riesgo (0-100)
	SuspiciousFlags []string `bson:"suspicious_flags,omitempty"` // Banderas de actividad sospechosa
	IsCompromised   bool     `bson:"is_compromised,omitempty"`   // Si se sospecha que está comprometido

	// Configuración y permisos
	Scopes      []string `bson:"scopes,omitempty"`      // Permisos del token
	Permissions []string `bson:"permissions,omitempty"` // Permisos específicos

	// Metadata adicional
	JTI         string                 `bson:"jti,omitempty"`         // JWT ID (si es JWT)
	Claims      map[string]interface{} `bson:"claims,omitempty"`      // Claims adicionales
	Fingerprint string                 `bson:"fingerprint,omitempty"` // Huella digital del dispositivo
	Notes       string                 `bson:"notes,omitempty"`       // Notas adicionales
}

// Log de actividad de tokens
type TokenActivityLog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	TokenID primitive.ObjectID `bson:"token_id"`

	// Identificación
	UserID    primitive.ObjectID  `bson:"user_id"`
	SessionID *primitive.ObjectID `bson:"session_id,omitempty"`

	Activity  string    `bson:"activity"` // created, used, refreshed, revoked, expired, validated
	Timestamp time.Time `bson:"timestamp"`

	// Información contextual
	ApplicationID  string     `bson:"application_id,omitempty"`
	ApplicationURL string     `bson:"application_url,omitempty"`
	DeviceInfo     DeviceInfo `bson:"device_info,omitempty"`

	// Detalles de la actividad
	Details TokenActivityDetails `bson:"details,omitempty"`

	// Relaciones
	AuthLogID     *primitive.ObjectID `bson:"auth_log_id,omitempty"`     // Log de auth relacionado
	NewTokenID    *primitive.ObjectID `bson:"new_token_id,omitempty"`    // Nuevo token generado (en refresh)
	ParentTokenID *primitive.ObjectID `bson:"parent_token_id,omitempty"` // Token padre usado

	// Información de resultado
	Success        bool   `bson:"success"`
	ErrorCode      string `bson:"error_code,omitempty"`
	ErrorMessage   string `bson:"error_message,omitempty"`
	ProcessingTime int    `bson:"processing_time,omitempty"` // ms
}

// Detalles específicos de actividad de tokens
type TokenActivityDetails struct {
	// Para creación de tokens
	GeneratedBy      string `bson:"generated_by,omitempty"`       // auth_attempt, refresh, etc.
	InitialExpiresIn int64  `bson:"initial_expires_in,omitempty"` // Duración inicial en segundos

	// Para uso de tokens
	EndpointAccessed string   `bson:"endpoint_accessed,omitempty"`
	HTTPMethod       string   `bson:"http_method,omitempty"`
	ResponseStatus   int      `bson:"response_status,omitempty"`
	ClaimsValidated  []string `bson:"claims_validated,omitempty"`

	// Para refresh de tokens
	RefreshReason string     `bson:"refresh_reason,omitempty"` // expired, proactive, security
	OldExpiresAt  *time.Time `bson:"old_expires_at,omitempty"`
	NewExpiresAt  *time.Time `bson:"new_expires_at,omitempty"`
	RefreshCount  int        `bson:"refresh_count,omitempty"`

	// Para revocación
	RevokedBy        string `bson:"revoked_by,omitempty"` // user, admin, system, application
	RevocationReason string `bson:"revocation_reason,omitempty"`
	CascadeRevoked   bool   `bson:"cascade_revoked,omitempty"` // Si revocó tokens relacionados

	// Para validación
	ValidationResult string   `bson:"validation_result,omitempty"` // valid, expired, invalid, revoked
	ValidationTime   int      `bson:"validation_time,omitempty"`   // ms
	ClaimsChecked    []string `bson:"claims_checked,omitempty"`

	// Información de seguridad
	IPChanged         bool `bson:"ip_changed,omitempty"`
	LocationChanged   bool `bson:"location_changed,omitempty"`
	DeviceChanged     bool `bson:"device_changed,omitempty"`
	SuspiciousPattern bool `bson:"suspicious_pattern,omitempty"`
	RiskScoreChange   int  `bson:"risk_score_change,omitempty"` // Cambio en puntuación de riesgo

	// Información adicional
	RequestID string `bson:"request_id,omitempty"` // ID único de la request
	TraceID   string `bson:"trace_id,omitempty"`   // ID de trace distribuido
	SpanID    string `bson:"span_id,omitempty"`    // ID del span
	UserAgent string `bson:"user_agent,omitempty"`
	Referer   string `bson:"referer,omitempty"`
}

// Estadísticas de tokens (para reporting y análisis)
type TokenStats struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Date   time.Time          `bson:"date"` // Fecha del reporte (día)

	// Contadores generales
	TotalTokensIssued   int `bson:"total_tokens_issued"`
	AccessTokensIssued  int `bson:"access_tokens_issued"`
	RefreshTokensIssued int `bson:"refresh_tokens_issued"`
	TokensRevoked       int `bson:"tokens_revoked"`
	TokensExpired       int `bson:"tokens_expired"`

	// Contadores de actividad
	TokenUsages       int `bson:"token_usages"`
	TokenRefreshes    int `bson:"token_refreshes"`
	TokenValidations  int `bson:"token_validations"`
	FailedValidations int `bson:"failed_validations"`

	// Información de aplicaciones
	ApplicationsAccessed []string `bson:"applications_accessed"`
	UniqueApplications   int      `bson:"unique_applications"`

	// Información de dispositivos
	UniqueDevices   int `bson:"unique_devices"`
	UniqueIPs       int `bson:"unique_ips"`
	UniqueLocations int `bson:"unique_locations"`

	// Información de seguridad
	SuspiciousActivities int `bson:"suspicious_activities"`
	SecurityWarnings     int `bson:"security_warnings"`
	CompromisedTokens    int `bson:"compromised_tokens"`
	HighRiskActivities   int `bson:"high_risk_activities"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
