package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrganicUnit struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null;unique" validate:"required" json:"name"`
	Acronym     string     `gorm:"type:varchar(20);unique" validate:"required" json:"acronym"`
	Brand       *string    `gorm:"type:varchar(100)" json:"brand"`
	Description *string    `gorm:"type:text" json:"description"`
	ParentID    *uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:now()" json:"updated_at"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	Parent        *OrganicUnit  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children      []OrganicUnit `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	DeletedByUser *User         `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	Users         []User        `gorm:"foreignKey:OrganicUnitID" json:"users,omitempty"`
}
