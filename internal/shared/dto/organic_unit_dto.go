package dto

type CreateOrganicUnitDTO struct {
	Name        string `json:"name" validate:"required"`
	Acronym     string `json:"acronym"`
	Brand       string `json:"brand"`
	Level       string `json:"level"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"` // opcional
}

type UpdateOrganicUnitDTO struct {
	Name        string `json:"name" validate:"required"`
	Acronym     string `json:"acronym"`
	Brand       string `json:"brand"`
	Level       string `json:"level"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"`
}
