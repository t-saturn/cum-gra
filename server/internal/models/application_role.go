package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ApplicationRole struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name          string         `gorm:"type:varchar(100);not null"`
	Description   string         `gorm:"type:text"`
	ApplicationID uuid.UUID      `gorm:"type:uuid;not null"`
	Permissions   pq.StringArray `gorm:"type:text[]"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`

	// Relaciones
	Application           Application            `gorm:"foreignKey:ApplicationID"`
	UserApplicationRoles  []UserApplicationRole  `gorm:"foreignKey:ApplicationRoleID"`
	ModuleRolePermissions []ModuleRolePermission `gorm:"foreignKey:ApplicationRoleID"`
}

func (ar *ApplicationRole) BeforeCreate(tx *gorm.DB) error {
	if ar.ID == uuid.Nil {
		ar.ID = uuid.New()
	}
	return nil
}
