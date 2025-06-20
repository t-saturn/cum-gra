package dto

import "time"

type CreateUserModuleRestrictionDTO struct {
	UserID             string     `json:"user_id" validate:"required,uuid"`
	ModuleID           string     `json:"module_id" validate:"required,uuid"`
	ApplicationID      string     `json:"application_id" validate:"required,uuid"`
	RestrictionType    string     `json:"restriction_type" validate:"required,oneof=block_access limit_permission"`
	MaxPermissionLevel string     `json:"max_permission_level" validate:"omitempty,oneof=denied access"`
	Reason             string     `json:"reason"`
	ExpiresAt          *time.Time `json:"expires_at"`
	CreatedBy          string     `json:"created_by" validate:"required,uuid"`
	UpdatedBy          string     `json:"updated_by" validate:"required,uuid"`
}

type UpdateUserModuleRestrictionDTO struct {
	RestrictionType    string     `json:"restriction_type" validate:"required,oneof=block_access limit_permission"`
	MaxPermissionLevel string     `json:"max_permission_level" validate:"omitempty,oneof=denied access"`
	Reason             string     `json:"reason"`
	ExpiresAt          *time.Time `json:"expires_at"`
	UpdatedBy          string     `json:"updated_by" validate:"required,uuid"`
}
