package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserApplicationRole struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID            uuid.UUID `gorm:"type:uuid;not null"`
	ApplicationID     uuid.UUID `gorm:"type:uuid;not null"`
	ApplicationRoleID uuid.UUID `gorm:"type:uuid;not null"`
	GrantedAt         time.Time `gorm:"autoCreateTime"`
	GrantedBy         uuid.UUID `gorm:"type:uuid;not null"`
	RevokedAt         *time.Time
	RevokedBy         *uuid.UUID `gorm:"type:uuid"`

	// Relaciones
	User            User            `gorm:"foreignKey:UserID"`
	Application     Application     `gorm:"foreignKey:ApplicationID"`
	ApplicationRole ApplicationRole `gorm:"foreignKey:ApplicationRoleID"`
	GrantedByUser   User            `gorm:"foreignKey:GrantedBy"`
	RevokedByUser   *User           `gorm:"foreignKey:RevokedBy"`
}

func (uar *UserApplicationRole) BeforeCreate(tx *gorm.DB) error {
	if uar.ID == uuid.Nil {
		uar.ID = uuid.New()
	}
	return nil
}
