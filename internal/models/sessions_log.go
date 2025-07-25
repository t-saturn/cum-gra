package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sesión de usuario
type Session struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SessionID string             `bson:"session_id"` // UUID único de la sesión
	UserID    primitive.ObjectID `bson:"user_id"`

	// Estado de la sesión
	Status       string     `bson:"status"`    // active, inactive, revoked, expired
	IsActive     bool       `bson:"is_active"` // Estado rápido de consulta
	CreatedAt    time.Time  `bson:"created_at"`
	LastActivity time.Time  `bson:"last_activity"`
	ExpiresAt    time.Time  `bson:"expires_at"`
	RevokedAt    *time.Time `bson:"revoked_at,omitempty"`

	// Información del dispositivo y ubicación
	DeviceInfo SessionDeviceInfo `bson:"device_info,omitempty"`

	// Métricas y estadísticas
	Metrics SessionMetrics `bson:"metrics,omitempty"`

	// Relaciones
	AuthAttemptID   *primitive.ObjectID  `bson:"auth_attempt_id,omitempty"`   // Intento que creó la sesión
	ParentSessionID *primitive.ObjectID  `bson:"parent_session_id,omitempty"` // Sesión padre (si aplica)
	ActiveTokens    []primitive.ObjectID `bson:"active_tokens,omitempty"`     // Tokens activos de esta sesión

	// Información de revocación
	RevokedBy        *primitive.ObjectID `bson:"revoked_by,omitempty"`        // Usuario/admin que revocó
	RevocationReason string              `bson:"revocation_reason,omitempty"` // Razón de revocación

	// Configuración de sesión
	MaxInactivity    time.Duration `bson:"max_inactivity,omitempty"` // Tiempo máximo de inactividad
	ExtendOnActivity bool          `bson:"extend_on_activity"`       // Si se extiende con actividad

	// Información de seguridad
	RiskScore        int      `bson:"risk_score,omitempty"`        // Puntuación de riesgo (0-100)
	SuspiciousFlags  []string `bson:"suspicious_flags,omitempty"`  // Banderas de actividad sospechosa
	SecurityWarnings []string `bson:"security_warnings,omitempty"` // Advertencias de seguridad

	// Metadata
	ApplicationID  string `bson:"application_id,omitempty"`  // Aplicación que inició la sesión
	ApplicationURL string `bson:"application_url,omitempty"` // URL de la aplicación
	Notes          string `bson:"notes,omitempty"`           // Notas adicionales
}

// Log de actividad de sesión (para auditoría detallada)
type SessionActivityLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SessionID primitive.ObjectID `bson:"session_id"`
	UserID    primitive.ObjectID `bson:"user_id"`

	Activity   string     `bson:"activity"` // created, accessed, refreshed, revoked, expired
	Timestamp  time.Time  `bson:"timestamp"`
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Detalles específicos de la actividad
	Details SessionActivityDetails `bson:"details,omitempty"`

	// Relaciones
	TokenID       *primitive.ObjectID `bson:"token_id,omitempty"`       // Token relacionado
	AuthLogID     *primitive.ObjectID `bson:"auth_log_id,omitempty"`    // Log de auth relacionado
	ApplicationID string              `bson:"application_id,omitempty"` // Aplicación que causó la actividad
}

// Detalles específicos de actividad de sesión
type SessionActivityDetails struct {
	// Para actividad de acceso
	EndpointAccessed string `bson:"endpoint_accessed,omitempty"`
	HTTPMethod       string `bson:"http_method,omitempty"`
	ResponseStatus   int    `bson:"response_status,omitempty"`

	// Para refreshes de token
	TokenType     string `bson:"token_type,omitempty"`
	OldTokenID    string `bson:"old_token_id,omitempty"`
	NewTokenID    string `bson:"new_token_id,omitempty"`
	RefreshReason string `bson:"refresh_reason,omitempty"`

	// Para revocación
	RevokedBy        string `bson:"revoked_by,omitempty"` // user, admin, system
	RevocationReason string `bson:"revocation_reason,omitempty"`

	// Para expiración
	ExpiredReason string `bson:"expired_reason,omitempty"` // timeout, manual, policy

	// Información adicional
	ProcessingTime  int    `bson:"processing_time,omitempty"` // ms
	ErrorCode       string `bson:"error_code,omitempty"`
	ErrorMessage    string `bson:"error_message,omitempty"`
	IPChanged       bool   `bson:"ip_changed,omitempty"`
	LocationChanged bool   `bson:"location_changed,omitempty"`
	DeviceChanged   bool   `bson:"device_changed,omitempty"`
}
