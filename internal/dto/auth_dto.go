package dto

type AuthVerifyRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	DNI      *string `json:"dni" validate:"omitempty,len=8"`
	Password string  `json:"password" validate:"required"`
}

type AuthVerifyResponse struct {
	UserID string `json:"user_id"`
}
