package dto

import (
	"time"

	"github.com/google/uuid"
)

type AuthSinginRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	DNI      *string `json:"dni" validate:"omitempty,len=8"`
	Password string  `json:"password" validate:"required"`
}

type AuthSinginResponse struct {
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"is_deleted"`
}

type ModuleDTO struct {
	ID        uuid.UUID  `json:"id"`
	Item      *string    `json:"item,omitempty"`
	Name      string     `json:"name"`
	Route     *string    `json:"route,omitempty"`
	Icon      *string    `json:"icon,omitempty"`
	ParentID  *uuid.UUID `json:"parent_id,omitempty"`
	SortOrder int        `json:"sort_order"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Parent   *SimpleModuleDTO  `json:"parent,omitempty"`
	Children []SimpleModuleDTO `json:"children,omitempty"`
}

type SimpleModuleDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AuthRoleRequest struct {
	UserID   string `json:"user_id"`
	ClientID string `json:"client_id"`
}

type AuthRoleResponse struct {
	RoleID  string      `json:"role_id"`
	Modules []ModuleDTO `json:"modules"`
}
