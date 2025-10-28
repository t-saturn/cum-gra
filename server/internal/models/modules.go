package models

import (
	"time"

	"github.com/google/uuid"
)

type Module struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Item          *string    `gorm:"type:varchar(100)" json:"item"`
	Name          string     `gorm:"type:varchar(100);not null" json:"name"`
	Route         *string    `gorm:"type:varchar(255)" json:"route"`
	Icon          *string    `gorm:"type:varchar(100)" json:"icon"`
	ParentID      *uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	ApplicationID *uuid.UUID `gorm:"type:uuid" json:"application_id"`
	SortOrder     int        `gorm:"default:0" json:"sort_order"`
	Status        string     `gorm:"type:module_status_enum;default:'active'" json:"status"`
	CreatedAt     time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:now()" json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeletedBy     *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	Parent                 *Module                 `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children               []Module                `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Application            *Application            `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	DeletedByUser          *User                   `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	ModuleRolePermissions  []ModuleRolePermission  `gorm:"foreignKey:ModuleID" json:"module_role_permissions,omitempty"`
	UserModuleRestrictions []UserModuleRestriction `gorm:"foreignKey:ModuleID" json:"user_module_restrictions,omitempty"`
}
