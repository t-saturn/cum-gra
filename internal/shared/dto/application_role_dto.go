package dto

type CreateApplicationRoleDTO struct {
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description"`
	ApplicationID string `json:"application_id" validate:"required"`
}

type UpdateApplicationRoleDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
