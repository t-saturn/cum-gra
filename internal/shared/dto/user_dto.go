package dto

type CreateUserDTO struct {
	Email                string `json:"email" validate:"required,email"`
	PasswordHash         string `json:"password_hash" validate:"required"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Phone                string `json:"phone"`
	DNI                  string `json:"dni" validate:"required,len=8"`
	StructuralPositionID string `json:"structural_position_id" validate:"required"`
	OrganicUnitID        string `json:"organic_unit_id" validate:"required"`
}

type UpdateUserDTO struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Phone                string `json:"phone"`
	StructuralPositionID string `json:"structural_position_id"`
	OrganicUnitID        string `json:"organic_unit_id"`
	Status               string `json:"status" validate:"required,oneof=active suspended deleted"`
}
