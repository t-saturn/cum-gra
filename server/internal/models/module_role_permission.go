package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionType string

const (
	PermissionDenied PermissionType = "denied"
	PermissionRead   PermissionType = "read"
	PermissionWrite  PermissionType = "write"
	PermissionAdmin  PermissionType = "admin"
)

type ModuleRolePermission struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ModuleID          uuid.UUID      `gorm:"type:uuid;not null"`
	ApplicationRoleID uuid.UUID      `gorm:"type:uuid;not null"`
	PermissionType    PermissionType `gorm:"type:varchar(50);not null"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`

	// Relaciones
	Module          Module          `gorm:"foreignKey:ModuleID"`
	ApplicationRole ApplicationRole `gorm:"foreignKey:ApplicationRoleID"`
}

func (mrp *ModuleRolePermission) BeforeCreate(tx *gorm.DB) error {
	if mrp.ID == uuid.Nil {
		mrp.ID = uuid.New()
	}
	return nil
}
