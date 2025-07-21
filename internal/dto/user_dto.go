package dto

// CreateUserDTO representa los datos necesarios para registrar un nuevo usuario.
type CreateUserDTO struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=22"`
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string `json:"phone" validate:"omitempty,e164"`
	DNI       string `json:"dni" validate:"required,len=8,numeric"`
}
