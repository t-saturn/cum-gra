package dto

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
