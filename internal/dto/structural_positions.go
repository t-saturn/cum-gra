package dto

import "time"

type StructuralPositionItemDTO struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	Level       *int       `json:"level,omitempty"`
	Description *string    `json:"description,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	IsDeleted   bool       `json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedBy   *string    `json:"deleted_by,omitempty"`

	UsersCount int64 `json:"users_count"`
}

type StructuralPositionsListResponse struct {
	Data     []StructuralPositionItemDTO `json:"data"`
	Total    int64                       `json:"total"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
}
