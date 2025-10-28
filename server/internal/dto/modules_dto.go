package dto

import (
	"time"
)

type ModuleWithAppDTO struct {
	ID            string     `json:"id"`
	Item          *string    `json:"item,omitempty"`
	Name          string     `json:"name"`
	Route         *string    `json:"route,omitempty"`
	Icon          *string    `json:"icon,omitempty"`
	ParentID      *string    `json:"parent_id,omitempty"`
	ApplicationID *string    `json:"application_id,omitempty"`
	SortOrder     int        `json:"sort_order"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	DeletedBy     *string    `json:"deleted_by,omitempty"`

	ApplicationName     *string `json:"application_name,omitempty"`
	ApplicationClientID *string `json:"application_client_id,omitempty"`
	UsersCount          int64   `json:"users_count"`
}

type ModulesListResponse struct {
	Data     []ModuleWithAppDTO `json:"data"`
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
}

type ModulesStatsResponse struct {
	TotalModules   int64 `json:"total_modules"`
	ActiveModules  int64 `json:"active_modules"`
	DeletedModules int64 `json:"deleted_modules"`
	TotalUsers     int64 `json:"total_users"`
}
