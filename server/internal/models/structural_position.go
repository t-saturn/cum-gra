package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StructuralPosition struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Code        string    `gorm:"type:varchar(50);unique;not null"`
	Level       string    `gorm:"type:varchar(50)"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	// Relaciones inversas
	Users []User `gorm:"foreignKey:StructuralPositionID"`
}

func (sp *StructuralPosition) BeforeCreate(tx *gorm.DB) error {
	if sp.ID == uuid.Nil {
		sp.ID = uuid.New()
	}
	return nil
}
