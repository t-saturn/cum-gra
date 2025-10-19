package dto

import "time"

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

	UsersCount int64 `json:"users_count"`
}

type OrganicUnitsListResponse struct {
	Data     []OrganicUnitItemDTO `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
}
