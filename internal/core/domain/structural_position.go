package domain

import (
	"time"

	"github.com/google/uuid"
)

type StructuralPosition struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Code        string     `gorm:"type:varchar(50);unique;not null" json:"code"`
	Level       *int       `gorm:"type:integer;unique" json:"level"`
	Description *string    `gorm:"type:text" json:"description"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:now()" json:"updated_at"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	DeletedByUser *User  `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	Users         []User `gorm:"foreignKey:StructuralPositionID" json:"users,omitempty"`
}
