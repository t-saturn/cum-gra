package dto

import (
	"time"
)

type ModuleRolePermissionDTO struct {
	ID                string     `json:"id"`
	ModuleID          string     `json:"module_id"`
	ApplicationRoleID string     `json:"application_role_id"`
	PermissionType    string     `json:"permission_type"`
	CreatedAt         time.Time  `json:"created_at"`
	IsDeleted         bool       `json:"is_deleted"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	DeletedBy         *string    `json:"deleted_by,omitempty"`
	
	// Información relacionada
	ModuleName          string  `json:"module_name"`
	ModuleRoute         *string `json:"module_route,omitempty"`
	RoleName            string  `json:"role_name"`
	ApplicationName     string  `json:"application_name"`
	ApplicationClientID string  `json:"application_client_id"`
}

type ModuleRolePermissionsListResponse struct {
	Data     []ModuleRolePermissionDTO `json:"data"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
}

type CreateModuleRolePermissionRequest struct {
	ModuleID          string `json:"module_id" validate:"required,uuid"`
	ApplicationRoleID string `json:"application_role_id" validate:"required,uuid"`
	PermissionType    string `json:"permission_type" validate:"required,oneof=read write execute delete admin"`
}

type UpdateModuleRolePermissionRequest struct {
	PermissionType *string `json:"permission_type" validate:"omitempty,oneof=read write execute delete admin"`
}

type ModuleRolePermissionsStatsResponse struct {
	TotalPermissions    int64 `json:"total_permissions"`
	ActivePermissions   int64 `json:"active_permissions"`
	DeletedPermissions  int64 `json:"deleted_permissions"`
	UniqueModules       int64 `json:"unique_modules"`
	UniqueRoles         int64 `json:"unique_roles"`
	PermissionsByType   map[string]int64 `json:"permissions_by_type"`
}

// DTO para asignar permisos masivos
type BulkAssignPermissionsRequest struct {
	ApplicationRoleID string   `json:"application_role_id" validate:"required,uuid"`
	ModuleIDs         []string `json:"module_ids" validate:"required,min=1,dive,uuid"`
	PermissionType    string   `json:"permission_type" validate:"required,oneof=read write execute delete admin"`
}

type BulkAssignPermissionsResponse struct {
	Created int                         `json:"created"`
	Skipped int                         `json:"skipped"`
	Failed  int                         `json:"failed"`
	Details []ModuleRolePermissionDTO   `json:"details"`
}