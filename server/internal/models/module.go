package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModuleStatus string

const (
	ModuleActive   ModuleStatus = "active"
	ModuleInactive ModuleStatus = "inactive"
)

type Module struct {
	ID            uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Item          string       `gorm:"type:varchar(100)"`
	Name          string       `gorm:"type:varchar(100);not null"`
	Label         string       `gorm:"type:varchar(100)"`
	Route         string       `gorm:"type:varchar(255)"`
	Icon          string       `gorm:"type:varchar(100)"`
	ParentID      *uuid.UUID   `gorm:"type:uuid"`
	ApplicationID uuid.UUID    `gorm:"type:uuid;not null"`
	SortOrder     int          `gorm:"default:0"`
	IsMenuItem    bool         `gorm:"default:true"`
	Status        ModuleStatus `gorm:"type:varchar(50);default:'active'"`
	CreatedAt     time.Time    `gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime"`

	// Relaciones
	Parent      *Module     `gorm:"foreignKey:ParentID"`
	Application Application `gorm:"foreignKey:ApplicationID"`

	// Relaciones inversas
	Children              []Module               `gorm:"foreignKey:ParentID"`
	ModuleRolePermissions []ModuleRolePermission `gorm:"foreignKey:ModuleID"`
}

func (m *Module) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
