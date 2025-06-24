package dto

import "github.com/google/uuid"

type CreateUserApplicationRoleDTO struct {
	UserID            uuid.UUID `json:"user_id"`
	ApplicationID     uuid.UUID `json:"application_id"`
	ApplicationRoleID uuid.UUID `json:"application_role_id"`
	GrantedBy         uuid.UUID `json:"granted_by"`
}
