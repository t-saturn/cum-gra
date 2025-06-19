package domain

import (
	"time"

	"github.com/google/uuid"
)

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
