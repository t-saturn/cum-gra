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

type ModulesAccess struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuthRoleResponse struct {
	RoleID  string          `json:"role_id"`
	Modules []ModulesAccess `json:"modules"`
}
