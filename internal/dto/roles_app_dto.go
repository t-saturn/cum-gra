package dto

import "github.com/google/uuid"

type AppMinimalDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ClientID string    `json:"client_id"`
}

type RoleMinimalDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ModuleMinimalDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Icon *string   `json:"icon"`
}

type RolesAppResponse struct {
	Apps    []AppMinimalDTO    `json:"apps"`
	Roles   []RoleMinimalDTO   `json:"roles"`
	Modules []ModuleMinimalDTO `json:"modules"`
}
