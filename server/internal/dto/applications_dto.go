package dto

import (
	"time"

)

type AdminUserDTO struct {
	FullName string `json:"full_name"`
	DNI      string `json:"dni"`
	Email    string `json:"email"`
}

type ApplicationDTO struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	ClientID    string          `json:"client_id"`
	Domain      string          `json:"domain"`
	Logo        *string         `json:"logo"`
	Description *string         `json:"description"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	IsDeleted   bool            `json:"is_deleted"`
	DeletedAt   *time.Time      `json:"deleted_at"`
	DeletedBy   *string         `json:"deleted_by"`
	Admins      []AdminUserDTO  `json:"admins"`
	UsersCount  int64           `json:"users_count"`
}

type ApplicationsListResponse struct {
	Data     []ApplicationDTO `json:"data"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}

type CreateApplicationRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	ClientID    string  `json:"client_id" validate:"required,min=3,max=255"`
	ClientSecret string `json:"client_secret" validate:"required,min=8"`
	Domain      string  `json:"domain" validate:"required,url"`
	Logo        *string `json:"logo" validate:"omitempty,url"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Status      string  `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateApplicationRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=100"`
	ClientID    *string `json:"client_id" validate:"omitempty,min=3,max=255"`
	ClientSecret *string `json:"client_secret" validate:"omitempty,min=8"`
	Domain      *string `json:"domain" validate:"omitempty,url"`
	Logo        *string `json:"logo" validate:"omitempty,url"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Status      *string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type ApplicationStatsResponse struct {
	TotalApplications int64 `json:"total_applications"`
	ActiveApplications int64 `json:"active_applications"`
	InactiveApplications int64 `json:"inactive_applications"`
	DeletedApplications int64 `json:"deleted_applications"`
}