package dto

import "time"

type UserApplicationRoleDTO struct {
	ID                string     `json:"id"`
	UserID            string     `json:"user_id"`
	ApplicationID     string     `json:"application_id"`
	ApplicationRoleID string     `json:"application_role_id"`
	GrantedAt         time.Time  `json:"granted_at"`
	GrantedBy         string     `json:"granted_by"`
	RevokedAt         *time.Time `json:"revoked_at,omitempty"`
	RevokedBy         *string    `json:"revoked_by,omitempty"`
	IsDeleted         bool       `json:"is_deleted"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	DeletedBy         *string    `json:"deleted_by,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	// Información relacionada
	UserEmail           string  `json:"user_email"`
	UserFullName        string  `json:"user_full_name"`
	ApplicationName     string  `json:"application_name"`
	ApplicationClientID string  `json:"application_client_id"`
	RoleName            string  `json:"role_name"`
	GrantedByEmail      string  `json:"granted_by_email"`
	RevokedByEmail      *string `json:"revoked_by_email,omitempty"`
}

type UserApplicationRolesListResponse struct {
	Data     []UserApplicationRoleDTO `json:"data"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
}

type CreateUserApplicationRoleRequest struct {
	UserID            string `json:"user_id" validate:"required,uuid"`
	ApplicationID     string `json:"application_id" validate:"required,uuid"`
	ApplicationRoleID string `json:"application_role_id" validate:"required,uuid"`
}

type RevokeUserApplicationRoleRequest struct {
	Reason *string `json:"reason" validate:"omitempty,max=500"`
}

type UserApplicationRolesStatsDTO struct {
	TotalAssignments   int64 `json:"total_assignments"`
	ActiveAssignments  int64 `json:"active_assignments"`
	RevokedAssignments int64 `json:"revoked_assignments"`
	DeletedAssignments int64 `json:"deleted_assignments"`
	UsersWithRoles     int64 `json:"users_with_roles"`
}

// DTO para asignar roles masivos a un usuario
type BulkAssignRolesToUserRequest struct {
	UserID        string   `json:"user_id" validate:"required,uuid"`
	ApplicationID string   `json:"application_id" validate:"required,uuid"`
	RoleIDs       []string `json:"role_ids" validate:"required,min=1,dive,uuid"`
}

type BulkAssignRolesToUserResponse struct {
	Created int                        `json:"created"`
	Skipped int                        `json:"skipped"`
	Failed  int                        `json:"failed"`
	Details []UserApplicationRoleDTO   `json:"details"`
}

// DTO para asignar un rol a múltiples usuarios
type BulkAssignRoleToUsersRequest struct {
	UserIDs           []string `json:"user_ids" validate:"required,min=1,dive,uuid"`
	ApplicationID     string   `json:"application_id" validate:"required,uuid"`
	ApplicationRoleID string   `json:"application_role_id" validate:"required,uuid"`
}

type BulkAssignRoleToUsersResponse struct {
	Created int                        `json:"created"`
	Skipped int                        `json:"skipped"`
	Failed  int                        `json:"failed"`
	Details []UserApplicationRoleDTO   `json:"details"`
}