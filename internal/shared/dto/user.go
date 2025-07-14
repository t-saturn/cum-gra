package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Email                string     `json:"email" validate:"required,email"`
	Password             string     `json:"password" validate:"required,min=8"`
	FirstName            *string    `json:"first_name"`
	LastName             *string    `json:"last_name"`
	Phone                *string    `json:"phone"`
	DNI                  string     `json:"dni" validate:"required,len=8"`
	StructuralPositionID *uuid.UUID `json:"structural_position_id"`
	OrganicUnitID        *uuid.UUID `json:"organic_unit_id"`
}
type UpdateUserDTO struct {
	Email                *string    `json:"email" validate:"omitempty,email"`
	Password             *string    `json:"password" validate:"omitempty,min=8"`
	FirstName            *string    `json:"first_name"`
	LastName             *string    `json:"last_name"`
	Phone                *string    `json:"phone"`
	DNI                  *string    `json:"dni" validate:"omitempty,len=8"`
	StructuralPositionID *uuid.UUID `json:"structural_position_id"`
	OrganicUnitID        *uuid.UUID `json:"organic_unit_id"`
	Status               *string    `json:"status" validate:"omitempty,oneof=active inactive suspended blocked"`
}
type UserResponseDTO struct {
	ID                   uuid.UUID  `json:"id"`
	Email                string     `json:"email"`
	FirstName            *string    `json:"first_name,omitempty"`
	LastName             *string    `json:"last_name,omitempty"`
	Phone                *string    `json:"phone,omitempty"`
	DNI                  string     `json:"dni"`
	EmailVerified        bool       `json:"email_verified"`
	PhoneVerified        bool       `json:"phone_verified"`
	TwoFactorEnabled     bool       `json:"two_factor_enabled"`
	Status               string     `json:"status"`
	StructuralPositionID *uuid.UUID `json:"structural_position_id,omitempty"`
	OrganicUnitID        *uuid.UUID `json:"organic_unit_id,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type CreateStructuralPositionDTO struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Code        string  `json:"code" validate:"required,min=3,max=10"`
	Level       *int    `json:"level" validate:"required,min=1"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
}
type UpdateStructuralPositionDTO struct {
	Name        *string `json:"name,omitempty"`
	Code        *string `json:"code,omitempty"`
	Level       *int    `json:"level,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}
type StructuralPositionResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Level       *int      `json:"level,omitempty"`
	Description *string   `json:"description,omitempty"`
	IsActive    bool      `json:"is_active"`
}

type CreateOrganicUnitDTO struct {
	Name        string     `json:"name" validate:"required"`
	Acronym     string     `json:"acronym" validate:"required"`
	Brand       *string    `json:"brand,omitempty"`
	Description *string    `json:"description,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
}
type UpdateOrganicUnitDTO struct {
	Name        *string    `json:"name,omitempty"`
	Acronym     *string    `json:"acronym,omitempty"`
	Brand       *string    `json:"brand,omitempty"`
	Description *string    `json:"description,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
}
type OrganicUnitResponseDTO struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Acronym     string     `json:"acronym"`
	Brand       *string    `json:"brand,omitempty"`
	Description *string    `json:"description,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
