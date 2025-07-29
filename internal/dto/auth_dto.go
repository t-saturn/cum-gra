package dto

// TokenValidationRequest representa el JSON recibido con el token.
type TokenValidationRequestDTO struct {
	Token string `json:"token" validate:"required"`
}

type TokenValidationResponseDTO struct {
	Valid     bool   `json:"valid"`
	Message   string `json:"message,omitempty"`
	Subject   string `json:"subject,omitempty"`
	IssuedAt  string `json:"issued_at,omitempty"`  // formato ISO8601
	ExpiresAt string `json:"expires_at,omitempty"` // formato ISO8601
	ExpiresIn int64  `json:"expires_in,omitempty"` // en segundos
}

type AuthVerifyRequestDTO struct {
	Email         *string       `json:"email,omitempty" validate:"omitempty,email"`
	DNI           *string       `json:"dni,omitempty" validate:"omitempty,len=8,numeric"`
	Password      string        `json:"password" validate:"required"`
	ApplicationID string        `json:"application_id" validate:"required,uuid4"`
	DeviceInfo    DeviceInfoDTO `json:"device_info" validate:"required"`
	CaptchaToken  *string       `json:"captcha_token,omitempty"`
}

type AuthVerifyResponseDTO struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
