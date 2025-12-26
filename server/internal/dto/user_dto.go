// internal/dto/user_dto.go
package dto

import "time"

// DTOs existentes
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
	Ubigeo             *SimpleUbigeoDTO             `json:"ubigeo,omitempty"`
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

// Nuevos DTOs para CRUD
type SimpleUbigeoDTO struct {
	ID         string `json:"id"`
	UbigeoCode string `json:"ubigeo_code"`
	Department string `json:"department"`
	Province   string `json:"province"`
	District   string `json:"district"`
}

type UserDetailDTO struct {
	ID                   string                       `json:"id"`
	Email                string                       `json:"email"`
	DNI                  string                       `json:"dni"`
	Status               string                       `json:"status"`
	FirstName            *string                      `json:"first_name,omitempty"`
	LastName             *string                      `json:"last_name,omitempty"`
	Phone                *string                      `json:"phone,omitempty"`
	CodEmpSGD            *string                      `json:"cod_emp_sgd,omitempty"`
	StructuralPositionID *string                      `json:"structural_position_id,omitempty"`
	OrganicUnitID        *string                      `json:"organic_unit_id,omitempty"`
	UbigeoID             *string                      `json:"ubigeo_id,omitempty"`
	CreatedAt            time.Time                    `json:"created_at"`
	UpdatedAt            time.Time                    `json:"updated_at"`
	IsDeleted            bool                         `json:"is_deleted"`
	DeletedAt            *time.Time                   `json:"deleted_at,omitempty"`
	DeletedBy            *string                      `json:"deleted_by,omitempty"`
	StructuralPosition   *SimpleStructuralPositionDTO `json:"structural_position,omitempty"`
	OrganicUnit          *SimpleOrganicUnitDTO        `json:"organic_unit,omitempty"`
	Ubigeo               *SimpleUbigeoDTO             `json:"ubigeo,omitempty"`
}

type CreateUserRequest struct {
	Email                string  `json:"email" validate:"required,email"`
	DNI                  string  `json:"dni" validate:"required,len=8,numeric"`
	Password             string  `json:"password" validate:"required,min=4,max=50"` // Ahora requerido
	FirstName            string  `json:"first_name" validate:"required,min=2,max=100"`
	LastName             string  `json:"last_name" validate:"required,min=2,max=100"`
	Phone                *string `json:"phone" validate:"omitempty,max=20"`
	Status               *string `json:"status" validate:"omitempty,oneof=active suspended inactive"`
	CodEmpSGD            *string `json:"cod_emp_sgd" validate:"omitempty,max=5"`
	StructuralPositionID *string `json:"structural_position_id" validate:"omitempty"`
	OrganicUnitID        *string `json:"organic_unit_id" validate:"omitempty"`
	UbigeoID             *string `json:"ubigeo_id" validate:"omitempty"`
}

type UpdateUserRequest struct {
	Email                *string `json:"email" validate:"omitempty,email"`
	DNI                  *string `json:"dni" validate:"omitempty,len=8,numeric"`
	FirstName            *string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName             *string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Phone                *string `json:"phone" validate:"omitempty,max=20"`
	Status               *string `json:"status" validate:"omitempty,oneof=active suspended inactive"`
	CodEmpSGD            *string `json:"cod_emp_sgd" validate:"omitempty,max=5"`
	StructuralPositionID *string `json:"structural_position_id" validate:"omitempty"`
	OrganicUnitID        *string `json:"organic_unit_id" validate:"omitempty"`
	UbigeoID             *string `json:"ubigeo_id" validate:"omitempty"`
}

// Para carga masiva
type BulkCreateUsersRequest struct {
	Users []CreateUserRequest `json:"users" validate:"required,dive"`
}

type BulkCreateUsersResponse struct {
	SuccessCount int                      `json:"success_count"`
	FailureCount int                      `json:"failure_count"`
	Results      []BulkCreateUserResult   `json:"results"`
}

type BulkCreateUserResult struct {
	DNI     string          `json:"dni"`
	Email   string          `json:"email"`
	Success bool            `json:"success"`
	User    *UserDetailDTO  `json:"user,omitempty"`
	Error   string          `json:"error,omitempty"`
}

type UserSelectDTO struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	DNI       string  `json:"dni"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Status    string  `json:"status"`
}