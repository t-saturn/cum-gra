package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Intento de autenticación
type AuthAttempt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Method        string             `bson:"method"` // "credentials" | "token"
	Status        string             `bson:"status"` // "pending", "success", "failed", "expired"
	ApplicationID string             `bson:"application_id"`

	// Campos para method = "credentials"
	Email        string `bson:"email,omitempty"`
	PasswordHash string `bson:"password_hash,omitempty"`

	// Campos para method = "token"
	Token     string `bson:"token,omitempty"`
	TokenType string `bson:"token_type,omitempty"` // "access" | "refresh"

	// Información del dispositivo
	DeviceInfo DeviceInfo `bson:"device_info,omitempty"`

	// Timestamps
	CreatedAt   time.Time  `bson:"created_at"`
	ValidatedAt *time.Time `bson:"validated_at,omitempty"`
	ExpiresAt   *time.Time `bson:"expires_at,omitempty"`

	// Respuesta de validación
	ValidationResponse *ValidationResponse `bson:"validation_response,omitempty"`

	// Relaciones
	UserID        *primitive.ObjectID `bson:"user_id,omitempty"`         // Se establece después de validación exitosa
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`      // Sesión asociada (si existe)
	CaptchaLogID  *primitive.ObjectID `bson:"captcha_log_id,omitempty"`  // CAPTCHA validado
	ParentTokenID *primitive.ObjectID `bson:"parent_token_id,omitempty"` // Token padre (para refresh)

	// Información adicional de seguridad
	FailedAttempts int        `bson:"failed_attempts,omitempty"` // Intentos fallidos previos
	LastFailedIP   string     `bson:"last_failed_ip,omitempty"`  // Última IP que falló
	IsBlocked      bool       `bson:"is_blocked,omitempty"`      // Si está bloqueado por seguridad
	BlockedUntil   *time.Time `bson:"blocked_until,omitempty"`   // Hasta cuándo está bloqueado
	RiskScore      int        `bson:"risk_score,omitempty"`      // Puntuación de riesgo (0-100)

	// Metadata
	ProcessingTime int    `bson:"processing_time,omitempty"` // Tiempo de procesamiento en ms
	ErrorCode      string `bson:"error_code,omitempty"`      // Código de error específico
	ErrorMessage   string `bson:"error_message,omitempty"`   // Mensaje de error
	Notes          string `bson:"notes,omitempty"`           // Notas adicionales
}

// Log completo de autenticación (para auditoría)
type AuthLog struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	UserID        primitive.ObjectID  `bson:"user_id,omitempty"`
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`
	AuthAttemptID *primitive.ObjectID `bson:"auth_attempt_id,omitempty"` // Referencia al intento
	CredentialID  string              `bson:"credential_id,omitempty"`
	TokenID       string              `bson:"token_id,omitempty"`

	Action         string `bson:"action,omitempty"` // login, token_validation, captcha, logout, etc.
	Success        bool   `bson:"success,omitempty"`
	ApplicationID  string `bson:"application_id,omitempty"`
	ApplicationURL string `bson:"application_url,omitempty"`

	Timestamp      time.Time `bson:"timestamp"`
	ProcessingTime int       `bson:"processing_time,omitempty"` // ms

	DeviceInfo DeviceInfo     `bson:"device_info,omitempty"`
	Details    AuthLogDetails `bson:"details,omitempty"`

	// Relaciones adicionales
	CaptchaLogID *primitive.ObjectID `bson:"captcha_log_id,omitempty"`
	TokenLogID   *primitive.ObjectID `bson:"token_log_id,omitempty"` // Referencia al token usado

	// Campos de auditoría
	RequestID string `bson:"request_id,omitempty"` // ID único de la request
	TraceID   string `bson:"trace_id,omitempty"`   // ID de trace distribuido
	SpanID    string `bson:"span_id,omitempty"`    // ID del span
}
