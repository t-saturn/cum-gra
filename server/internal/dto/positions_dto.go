package dto

import "time"

// DTOs existentes
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
	CodCarSGD   *string    `json:"cod_car_sgd,omitempty"`

	UsersCount int64 `json:"users_count"`
}

type StructuralPositionsListResponse struct {
	Data     []StructuralPositionItemDTO `json:"data"`
	Total    int64                       `json:"total"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
}

type StructuralPositionsStatsResponse struct {
	TotalPositions    int64 `json:"total_positions"`
	ActivePositions   int64 `json:"active_positions"`
	DeletedPositions  int64 `json:"deleted_positions"`
	AssignedEmployees int64 `json:"assigned_employees"`
}

// Nuevos DTOs para CRUD
type CreateStructuralPositionRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Code        string  `json:"code" validate:"required,min=2,max=50"`
	Level       *int    `json:"level" validate:"omitempty,min=1,max=10"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	IsActive    *bool   `json:"is_active"`
	CodCarSGD   *string `json:"cod_car_sgd" validate:"omitempty,max=4"`
}

type UpdateStructuralPositionRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=255"`
	Code        *string `json:"code" validate:"omitempty,min=2,max=50"`
	Level       *int    `json:"level" validate:"omitempty,min=1,max=10"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	IsActive    *bool   `json:"is_active"`
	CodCarSGD   *string `json:"cod_car_sgd" validate:"omitempty,max=4"`
}