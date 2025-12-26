package dto

type AuthSinginRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	DNI      *string `json:"dni" validate:"omitempty,len=8"`
	Password string  `json:"password" validate:"required"`
}

type AuthSinginResponse struct {
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"is_deleted"`
}

type AuthRoleRequest struct {
	ClientID string `json:"client_id" validate:"required"`
}

type AuthRoleResponse struct {
	RoleID   string      `json:"role_id"`
	RoleName string      `json:"role_name"`
	Modules  []ModuleDTO `json:"modules"`
}
