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

type RolesAppData struct {
	Apps    []AppMinimalDTO    `json:"apps"`
	Roles   []RoleMinimalDTO   `json:"roles"`
	Modules []ModuleMinimalDTO `json:"modules"`
}

type RolesAppResponse struct {
	Data     RolesAppData `json:"data"`
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
}

type RolesAppStatsResponse struct {
	TotalRoles    int64 `json:"total_roles"`
	ActiveRoles   int64 `json:"active_roles"`
	AdminRoles    int64 `json:"admin_roles"`
	AssignedUsers int64 `json:"assigned_users"`
}
