package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthVerifyRequest representa la estructura de entrada para verificar credenciales de autenticación.
type AuthVerifyRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	DNI      *string `json:"dni" validate:"omitempty,len=8"`
	Password string  `json:"password" validate:"required"`
}

// AuthVerifyResponse representa la respuesta al verificar credenciales exitosamente, incluyendo el ID del usuario.
type AuthVerifyResponse struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenValidationRequest representa el JSON recibido con el token.
type TokenValidationRequest struct {
	Token string `json:"token" validate:"required"`
}

// TokenValidationResponse representa la respuesta tras validar el token.
type TokenValidationResponse struct {
	Valid     bool   `json:"valid"`
	Message   string `json:"message,omitempty"`
	Subject   string `json:"subject,omitempty"`
	IssuedAt  string `json:"issued_at,omitempty"`  // formato ISO8601
	ExpiresAt string `json:"expires_at,omitempty"` // formato ISO8601
	ExpiresIn int64  `json:"expires_in,omitempty"` // en segundos
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
// VerifyCredentialsDTO representa los datos necesarios para verificar credenciales de un usuario.
type VerifyCredentialsDTO struct {
	Email         string        `json:"email" validate:"omitempty,email"`       // email o dni requerido
	DNI           string        `json:"dni" validate:"omitempty,len=8,numeric"` // email o dni requerido
	Password      string        `json:"password" validate:"required,min=6,max=64"`
	ApplicationID string        `json:"application_id" validate:"required,uuid4"`
	DeviceInfo    DeviceInfoDTO `json:"device_info" validate:"required,dive"`
	CaptchaToken  string        `json:"captcha_token,omitempty" validate:"omitempty"`
}

// DeviceInfoDTO representa los datos básicos del dispositivo usados para validación.
type DeviceInfoDTO struct {
	UserAgent   string `json:"user_agent" validate:"required"`
	IP          string `json:"ip" validate:"required,ip"` // se puede usar `ip|ipv4|ipv6` según necesidad
	DeviceID    string `json:"device_id" validate:"omitempty"`
	BrowserName string `json:"browser_name" validate:"omitempty"`
	OS          string `json:"os" validate:"omitempty"`
}

// VerifyCredentialsResponseDTO representa la respuesta al verificar credenciales del usuario.
type VerifyCredentialsResponseDTO struct {
	UserID        primitive.ObjectID `json:"user_id"`
	Status        string             `json:"status"`          // e.g. "success"
	ValidatedBy   string             `json:"validated_by"`    // e.g. "credentials"
	AuthAttemptID primitive.ObjectID `json:"auth_attempt_id"` // ID del intento registrado
}
