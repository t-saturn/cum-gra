package dto

import "time"

type CreateUserDTO struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=22"`
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string `json:"phone" validate:"omitempty,e164"`
	DNI       string `json:"dni" validate:"required,len=8,numeric"`
}

type SimpleStructuralPositionDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Level *int   `json:"level,omitempty"`
}

type UserListItemDTO struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	DNI       string  `json:"dni"`
	Status    string  `json:"status"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *string    `json:"deleted_by,omitempty"`

	OrganicUnit        *SimpleOrganicUnitDTO        `json:"organic_unit,omitempty"`
	StructuralPosition *SimpleStructuralPositionDTO `json:"structural_position,omitempty"`
}

type UsersListResponse struct {
	Data     []UserListItemDTO `json:"data"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

type UsersStatsResponse struct {
	TotalUsers        int64 `json:"total_users"`
	ActiveUsers       int64 `json:"active_users"`
	SuspendedUsers    int64 `json:"suspended_users"`
	NewUsersLastMonth int64 `json:"new_users_last_month"`
}
