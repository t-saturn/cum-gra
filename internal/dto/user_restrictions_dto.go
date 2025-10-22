package dto

type UserAppAssignmentDTO struct {
	App             AppMinimalDTO      `json:"app"`
	Role            *RoleMinimalDTO    `json:"role,omitempty"`
	Modules         []ModuleMinimalDTO `json:"modules"`
	ModulesRestrict []ModuleMinimalDTO `json:"modules_restrict"`
}

// Estructura por usuario
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
