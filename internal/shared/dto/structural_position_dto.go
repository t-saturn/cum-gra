package dto

type CreateStructuralPositionDTO struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Level       string `json:"level"`
	Description string `json:"description"`
}

type UpdateStructuralPositionDTO struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Level       string `json:"level"`
	Description string `json:"description"`
}
