package domain

import (
	"time"

	"github.com/google/uuid"
)

type Module struct {
	ID            uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Item          *string    `gorm:"type:varchar(100)" json:"item"`
	Name          string     `gorm:"type:varchar(100);not null; unique" json:"name"`
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

	// Relaciones
	Parent                 *Module                 `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children               []Module                `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Application            *Application            `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	DeletedByUser          *User                   `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	ModuleRolePermissions  []ModuleRolePermission  `gorm:"foreignKey:ModuleID" json:"module_role_permissions,omitempty"`
	UserModuleRestrictions []UserModuleRestriction `gorm:"foreignKey:ModuleID" json:"user_module_restrictions,omitempty"`
}

type ModuleRolePermission struct {
	ID                uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	ModuleID          uuid.UUID  `gorm:"type:uuid;not null" json:"module_id"`
	ApplicationRoleID uuid.UUID  `gorm:"type:uuid;not null" json:"application_role_id"`
	PermissionType    string     `gorm:"type:permission_type_enum;not null" json:"permission_type"`
	CreatedAt         time.Time  `gorm:"default:now()" json:"created_at"`
	IsDeleted         bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt         *time.Time `json:"deleted_at"`
	DeletedBy         *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	Module          *Module          `gorm:"foreignKey:ModuleID" json:"module,omitempty"`
	ApplicationRole *ApplicationRole `gorm:"foreignKey:ApplicationRoleID" json:"application_role,omitempty"`
	DeletedByUser   *User            `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
}

type UserModuleRestriction struct {
	ID                 uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	UserID             uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	ModuleID           uuid.UUID  `gorm:"type:uuid;not null" json:"module_id"`
	ApplicationID      uuid.UUID  `gorm:"type:uuid;not null" json:"application_id"`
	RestrictionType    string     `gorm:"type:restriction_type_enum;not null" json:"restriction_type"`
	MaxPermissionLevel *string    `gorm:"type:permission_level_enum" json:"max_permission_level"`
	Reason             *string    `gorm:"type:text" json:"reason"`
	ExpiresAt          *time.Time `json:"expires_at"`
	CreatedAt          time.Time  `gorm:"default:now()" json:"created_at"`
	CreatedBy          uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedAt          time.Time  `gorm:"default:now()" json:"updated_at"`
	UpdatedBy          *uuid.UUID `gorm:"type:uuid" json:"updated_by"`
	IsDeleted          bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt          *time.Time `json:"deleted_at"`
	DeletedBy          *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	User          *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Module        *Module      `gorm:"foreignKey:ModuleID" json:"module,omitempty"`
	Application   *Application `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	CreatedByUser *User        `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *User        `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
	DeletedByUser *User        `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
}
