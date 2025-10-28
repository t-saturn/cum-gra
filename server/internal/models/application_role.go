package models

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationRole struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string     `gorm:"type:varchar(100);not null" json:"name"`
	Description   *string    `gorm:"type:text" json:"description"`
	ApplicationID uuid.UUID  `gorm:"type:uuid;not null" json:"application_id"`
	CreatedAt     time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:now()" json:"updated_at"`
	IsDeleted     bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeletedBy     *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	Application           *Application           `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	DeletedByUser         *User                  `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	UserApplicationRoles  []UserApplicationRole  `gorm:"foreignKey:ApplicationRoleID" json:"user_application_roles,omitempty"`
	ModuleRolePermissions []ModuleRolePermission `gorm:"foreignKey:ApplicationRoleID" json:"module_role_permissions,omitempty"`
}
