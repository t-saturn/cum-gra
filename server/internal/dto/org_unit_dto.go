package dto

import "time"

// DTOs existentes
type OrganicUnitItemDTO struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Acronym     string     `json:"acronym"`
	Brand       *string    `json:"brand,omitempty"`
	Description *string    `json:"description,omitempty"`
	ParentID    *string    `json:"parent_id,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	IsDeleted   bool       `json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedBy   *string    `json:"deleted_by,omitempty"`
	CodDepSGD   *string    `json:"cod_dep_sgd,omitempty"`

	UsersCount int64 `json:"users_count"`
}

type OrganicUnitsListResponse struct {
	Data     []OrganicUnitItemDTO `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
}

type OrganicUnitsStatsResponse struct {
	TotalOrganicUnits   int64 `json:"total_organic_units"`
	ActiveOrganicUnits  int64 `json:"active_organic_units"`
	DeletedOrganicUnits int64 `json:"deleted_organic_units"`
	TotalEmployees      int64 `json:"total_employees"`
}

// Nuevos DTOs para CRUD
type SimpleOrganicUnitDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Acronym  string `json:"acronym"`
	ParentID *string `json:"parent_id,omitempty"`
}

type CreateOrganicUnitRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Acronym     string  `json:"acronym" validate:"required,min=2,max=20"`
	Brand       *string `json:"brand" validate:"omitempty,max=100"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	ParentID    *string `json:"parent_id" validate:"omitempty"`
	IsActive    *bool   `json:"is_active"`
	CodDepSGD   *string `json:"cod_dep_sgd" validate:"omitempty,max=5"`
}

type UpdateOrganicUnitRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=255"`
	Acronym     *string `json:"acronym" validate:"omitempty,min=2,max=20"`
	Brand       *string `json:"brand" validate:"omitempty,max=100"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	ParentID    *string `json:"parent_id" validate:"omitempty"`
	IsActive    *bool   `json:"is_active"`
	CodDepSGD   *string `json:"cod_dep_sgd" validate:"omitempty,max=5"`
}

type OrganicUnitSelectDTO struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Acronym  string  `json:"acronym"`
	ParentID *string `json:"parent_id,omitempty"`
	IsActive bool    `json:"is_active"`
}