package dto

type CreateModuleRolePermissionDTO struct {
	ModuleID          string `json:"module_id" validate:"required,uuid"`
	ApplicationRoleID string `json:"application_role_id" validate:"required,uuid"`
	PermissionType    string `json:"permission_type" validate:"required,oneof=denied access"`
}

type UpdateModuleRolePermissionDTO struct {
	PermissionType string `json:"permission_type" validate:"required,oneof=denied access"`
}
