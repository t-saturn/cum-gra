package dto

type SessionMeRequest struct {
	UserID string `json:"user_id" validate:"required,uuid"`
}

type SessionMeResponse struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	DNI       string  `json:"dni"`
	Status    string  `json:"status"`
	IsDeleted bool    `json:"is_deleted"`
}
