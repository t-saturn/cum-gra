package dto

import "github.com/google/uuid"

type UserBriefDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	DNI       string    `json:"dni"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
}

type RestrictedModuleDTO struct {
	ModuleID           uuid.UUID `json:"module_id"`
	ModuleName         string    `json:"module_name"`
	ModuleIcon         *string   `json:"module_icon,omitempty"`
	RestrictionType    string    `json:"restriction_type"`
	MaxPermissionLevel *string   `json:"max_permission_level,omitempty"`
}

type UserAppAccessDTO struct {
	Application       AppMinimalDTO         `json:"application"`
	Role              RoleMinimalDTO        `json:"role"`
	RestrictedModules []RestrictedModuleDTO `json:"restricted_modules"`
}

type UserWithAppsDTO struct {
	User         UserBriefDTO       `json:"user"`
	Applications []UserAppAccessDTO `json:"applications"`
}

type UsersAppAccessResponse struct {
	Data     []UserWithAppsDTO `json:"data"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}
