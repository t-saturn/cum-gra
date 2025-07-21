package dto

// AuthVerifyRequest representa la estructura de entrada para verificar credenciales de autenticación.
type AuthVerifyRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	DNI      *string `json:"dni" validate:"omitempty,len=8"`
	Password string  `json:"password" validate:"required"`
}

// AuthVerifyResponse representa la respuesta al verificar credenciales exitosamente, incluyendo el ID del usuario.
type AuthVerifyResponse struct {
	UserID string `json:"user_id"`
}
