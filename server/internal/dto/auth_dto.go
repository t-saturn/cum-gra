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
	RoleID   string            `json:"role_id"`
	RoleName string            `json:"role_name"`
	Modules  []ModuleWithPerms `json:"modules"`
}

type ModuleWithPerms struct {
	ID             string              `json:"id"`
	Item           *string             `json:"item"`
	Name           string              `json:"name"`
	Route          *string             `json:"route"`
	Icon           *string             `json:"icon"`
	ParentID       *string             `json:"parent_id"`
	SortOrder      int                 `json:"sort_order"`
	Status         string              `json:"status"`
	PermissionType string              `json:"permission_type"`
	Restriction    *ModuleRestriction  `json:"restriction,omitempty"`
	Children       []ModuleWithPerms   `json:"children,omitempty"`
}

type ModuleRestriction struct {
	RestrictionType    string  `json:"restriction_type"`
	MaxPermissionLevel *string `json:"max_permission_level,omitempty"`
	Reason             *string `json:"reason,omitempty"`
	ExpiresAt          *string `json:"expires_at,omitempty"`
}