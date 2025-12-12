package dto

import (
	"time"

	"github.com/google/uuid"
)

// DTOs existentes que ya tienes
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

type SimpleModuleDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ModuleDTO struct {
	ID        uuid.UUID        `json:"id"`
	Item      *string          `json:"item,omitempty"`
	Name      string           `json:"name"`
	Route     *string          `json:"route,omitempty"`
	Icon      *string          `json:"icon,omitempty"`
	ParentID  *uuid.UUID       `json:"parent_id,omitempty"`
	SortOrder int              `json:"sort_order"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Parent    *SimpleModuleDTO `json:"parent,omitempty"`
	Children  []SimpleModuleDTO `json:"children,omitempty"`
}

type CreateModuleRequest struct {
	Item          *string `json:"item" validate:"omitempty,max=100"`
	Name          string  `json:"name" validate:"required,min=3,max=100"`
	Route         *string `json:"route" validate:"omitempty,max=255"`
	Icon          *string `json:"icon" validate:"omitempty,max=100"`
	ParentID      *string `json:"parent_id" validate:"omitempty,uuid"`
	ApplicationID *string `json:"application_id" validate:"omitempty,uuid"`
	SortOrder     *int    `json:"sort_order" validate:"omitempty,min=0"`
	Status        *string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateModuleRequest struct {
	Item      *string `json:"item" validate:"omitempty,max=100"`
	Name      *string `json:"name" validate:"omitempty,min=3,max=100"`
	Route     *string `json:"route" validate:"omitempty,max=255"`
	Icon      *string `json:"icon" validate:"omitempty,max=100"`
	ParentID  *string `json:"parent_id" validate:"omitempty,uuid"`
	SortOrder *int    `json:"sort_order" validate:"omitempty,min=0"`
	Status    *string `json:"status" validate:"omitempty,oneof=active inactive"`
}