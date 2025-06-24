package dto

type CreateModuleDTO struct {
	Item          string `json:"item"`
	Name          string `json:"name" validate:"required"`
	Label         string `json:"label"`
	Route         string `json:"route"`
	Icon          string `json:"icon"`
	ParentID      string `json:"parent_id"`
	ApplicationID string `json:"application_id" validate:"required"`
	SortOrder     int    `json:"sort_order"`
	Status        string `json:"status" validate:"required,oneof=active inactive"`
}

type UpdateModuleDTO struct {
	Item          string `json:"item"`
	Name          string `json:"name" validate:"required"`
	Label         string `json:"label"`
	Route         string `json:"route"`
	Icon          string `json:"icon"`
	ParentID      string `json:"parent_id"`
	ApplicationID string `json:"application_id" validate:"required"`
	SortOrder     int    `json:"sort_order"`
	Status        string `json:"status" validate:"required,oneof=active inactive"`
}
