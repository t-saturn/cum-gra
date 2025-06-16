package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserPermission struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null"`
	ApplicationID uuid.UUID      `gorm:"type:uuid;not null"`
	Scopes        pq.StringArray `gorm:"type:text[]"`
	GrantedAt     time.Time      `gorm:"autoCreateTime"`
	GrantedBy     uuid.UUID      `gorm:"type:uuid;not null"`
	RevokedAt     *time.Time
	RevokedBy     *uuid.UUID `gorm:"type:uuid"`

	// Relaciones
	User          User        `gorm:"foreignKey:UserID"`
	Application   Application `gorm:"foreignKey:ApplicationID"`
	GrantedByUser User        `gorm:"foreignKey:GrantedBy"`
	RevokedByUser *User       `gorm:"foreignKey:RevokedBy"`
}

func (up *UserPermission) BeforeCreate(tx *gorm.DB) error {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return nil
}
