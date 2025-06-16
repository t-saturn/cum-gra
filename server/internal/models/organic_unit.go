package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganicUnit struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Acronym     string     `gorm:"type:varchar(20)"`
	Brand       string     `gorm:"type:varchar(100)"`
	Level       string     `gorm:"type:varchar(50)"`
	Description string     `gorm:"type:text"`
	ParentID    *uuid.UUID `gorm:"type:uuid"`
	IsActive    bool       `gorm:"default:true"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`

	// Relaciones
	Parent *OrganicUnit `gorm:"foreignKey:ParentID"`

	// Relaciones inversas
	Children []OrganicUnit `gorm:"foreignKey:ParentID"`
	Users    []User        `gorm:"foreignKey:OrganicUnitID"`
}

func (ou *OrganicUnit) BeforeCreate(tx *gorm.DB) error {
	if ou.ID == uuid.Nil {
		ou.ID = uuid.New()
	}
	return nil
}
