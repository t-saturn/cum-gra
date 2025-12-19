package dto

import (
	"time"

	"github.com/google/uuid"
)

type AppMinimalDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ClientID string    `json:"client_id"`
}

type RoleMinimalDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ModuleMinimalDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Icon *string   `json:"icon"`
}

type RoleAppModulesItemDTO struct {
	Role        RoleMinimalDTO     `json:"role"`
	App         AppMinimalDTO      `json:"app"`
	AppModules  []ModuleMinimalDTO `json:"app_modules"`
	RoleModules []ModuleMinimalDTO `json:"role_modules"`
}

type RolesAppResponseDTO struct {
	Data     []RoleAppModulesItemDTO `json:"data"`
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"page_size"`
}

type RolesAppStatsResponse struct {
	TotalRoles    int64 `json:"total_roles"`
	ActiveRoles   int64 `json:"active_roles"`
	AdminRoles    int64 `json:"admin_roles"`
	AssignedUsers int64 `json:"assigned_users"`
}

// DTOs para CRUD de ApplicationRole
type ApplicationRoleDTO struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Description   *string        `json:"description"`
	ApplicationID string         `json:"application_id"`
	Application   *AppMinimalDTO `json:"application,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	IsDeleted     bool           `json:"is_deleted"`
	DeletedAt     *time.Time     `json:"deleted_at"`
	DeletedBy     *string        `json:"deleted_by"`
	ModulesCount  int64          `json:"modules_count"`
	UsersCount    int64          `json:"users_count"`
}

type ApplicationRolesListResponse struct {
	Data     []ApplicationRoleDTO `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
}

type CreateApplicationRoleRequest struct {
	Name          string  `json:"name" validate:"required,min=3,max=100"`
	Description   *string `json:"description" validate:"omitempty,max=500"`
	ApplicationID string  `json:"application_id" validate:"required,uuid"`
}

type UpdateApplicationRoleRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}

type ApplicationRolesStatsResponse struct {
	TotalRoles       int64 `json:"total_roles"`
	ActiveRoles      int64 `json:"active_roles"`
	DeletedRoles     int64 `json:"deleted_roles"`
	RolesWithModules int64 `json:"roles_with_modules"`
	RolesWithUsers   int64 `json:"roles_with_users"`
}
