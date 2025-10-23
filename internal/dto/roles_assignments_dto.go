package dto

import "github.com/google/uuid"

type UserMinimalDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Email     string    `json:"email"`
	DNI       string    `json:"dni"`
}

type UserAppRoleDTO struct {
	Application AppMinimalDTO  `json:"application"`
	Role        RoleMinimalDTO `json:"role"`
}

type RoleAssignmentsDTO struct {
	User        UserMinimalDTO   `json:"user"`
	Assignments []UserAppRoleDTO `json:"assignments"`
}

type RolesAssigmentsResponseDTO struct {
	Data     []RoleAssignmentsDTO `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
}

type RolesAssigmentsResponseResponse struct {
	TotalUsers        int64 `json:"total_users"`
	AdminUsers        int64 `json:"admin_users"`
	UsersWithRoles    int64 `json:"users_with_roles"`
	UsersWithoutRoles int64 `json:"users_without_roles"`
}
