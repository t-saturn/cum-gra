package dto

import "time"

type AdminUserDTO struct {
	FullName string `json:"full_name"`
	DNI      string `json:"dni"`
	Email    string `json:"email"`
}

type ApplicationDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ClientID    string  `json:"client_id"`
	Domain      string  `json:"domain"`
	Logo        *string `json:"logo,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *string    `json:"deleted_by,omitempty"`

	Admins     []AdminUserDTO `json:"admins"`
	UsersCount int64          `json:"users_count"`
}

type ApplicationsListResponse struct {
	Data     []ApplicationDTO `json:"data"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}
