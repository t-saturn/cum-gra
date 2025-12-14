package dto

import "time"

// DTOs existentes
type UserAppAssignmentDTO struct {
	App             AppMinimalDTO      `json:"app"`
	Role            *RoleMinimalDTO    `json:"role,omitempty"`
	Modules         []ModuleMinimalDTO `json:"modules"`
	ModulesRestrict []ModuleMinimalDTO `json:"modules_restrict"`
}

type RoleRestrictDTO struct {
	User UserMinimalDTO         `json:"user"`
	Apps []UserAppAssignmentDTO `json:"apps"`
}

type RolesRestrictResponseDTO struct {
	Data     []RoleRestrictDTO `json:"data"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

type UserRestrictionsStatsDTO struct {
	TotalRestrictions   int64 `json:"total_restrictions"`
	ActiveRestrictions  int64 `json:"active_restrictions"`
	RestrictedUsers     int64 `json:"restricted_users"`
	DeletedRestrictions int64 `json:"deleted_restrictions"`
}

// Nuevos DTOs para CRUD
type UserModuleRestrictionDTO struct {
	ID                 string     `json:"id"`
	UserID             string     `json:"user_id"`
	ModuleID           string     `json:"module_id"`
	ApplicationID      string     `json:"application_id"`
	RestrictionType    string     `json:"restriction_type"`
	MaxPermissionLevel *string    `json:"max_permission_level,omitempty"`
	Reason             *string    `json:"reason,omitempty"`
	ExpiresAt          *time.Time `json:"expires_at,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	CreatedBy          string     `json:"created_by"`
	UpdatedAt          time.Time  `json:"updated_at"`
	UpdatedBy          *string    `json:"updated_by,omitempty"`
	IsDeleted          bool       `json:"is_deleted"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
	DeletedBy          *string    `json:"deleted_by,omitempty"`

	// Información relacionada
	UserEmail           string  `json:"user_email"`
	UserFullName        string  `json:"user_full_name"`
	ModuleName          string  `json:"module_name"`
	ModuleRoute         *string `json:"module_route,omitempty"`
	ApplicationName     string  `json:"application_name"`
	ApplicationClientID string  `json:"application_client_id"`
}

type UserModuleRestrictionsListResponse struct {
	Data     []UserModuleRestrictionDTO `json:"data"`
	Total    int64                      `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
}

type CreateUserModuleRestrictionRequest struct {
	UserID             string  `json:"user_id" validate:"required,uuid"`
	ModuleID           string  `json:"module_id" validate:"required,uuid"`
	ApplicationID      string  `json:"application_id" validate:"required,uuid"`
	RestrictionType    string  `json:"restriction_type" validate:"required,oneof=block limit read_only"`
	MaxPermissionLevel *string `json:"max_permission_level" validate:"omitempty,oneof=read write execute delete admin"`
	Reason             *string `json:"reason" validate:"omitempty,max=1000"`
	ExpiresAt          *string `json:"expires_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

type UpdateUserModuleRestrictionRequest struct {
	RestrictionType    *string `json:"restriction_type" validate:"omitempty,oneof=block limit read_only"`
	MaxPermissionLevel *string `json:"max_permission_level" validate:"omitempty,oneof=read write execute delete admin"`
	Reason             *string `json:"reason" validate:"omitempty,max=1000"`
	ExpiresAt          *string `json:"expires_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

// DTO para asignar restricciones masivas
type BulkCreateUserModuleRestrictionsRequest struct {
	UserID          string   `json:"user_id" validate:"required,uuid"`
	ApplicationID   string   `json:"application_id" validate:"required,uuid"`
	ModuleIDs       []string `json:"module_ids" validate:"required,min=1,dive,uuid"`
	RestrictionType string   `json:"restriction_type" validate:"required,oneof=block limit read_only"`
	Reason          *string  `json:"reason" validate:"omitempty,max=1000"`
}

type BulkCreateUserModuleRestrictionsResponse struct {
	Created int                          `json:"created"`
	Skipped int                          `json:"skipped"`
	Failed  int                          `json:"failed"`
	Details []UserModuleRestrictionDTO   `json:"details"`
}