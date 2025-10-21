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

type RolesAssignmentsResponseDTO struct {
	Assignments []RoleAssignmentsDTO `json:"assignments"`
}
