package dto

type CreateStructuralPositionDTO struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Level       *int   `json:"level" validate:"omitempty,min=1"`
	Description string `json:"description"`
}

type UpdateStructuralPositionDTO struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Level       *int   `json:"level" validate:"omitempty,min=1"`
	Description string `json:"description"`
}
