package dto

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
