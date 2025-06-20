package dto

type CreateApplicationRoleDTO struct {
	Name          string   `json:"name" validate:"required"`
	Description   string   `json:"description"`
	ApplicationID string   `json:"application_id" validate:"required"`
	Permissions   []string `json:"permissions"`
}

type UpdateApplicationRoleDTO struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}
