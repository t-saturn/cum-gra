package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganicUnit struct {
    ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    Name        string     `gorm:"type:varchar(255);not null;unique" json:"name"`
    Acronym     string     `gorm:"type:varchar(20);unique" json:"acronym"`
    Brand       *string    `gorm:"type:varchar(100)" json:"brand"`
    Description *string    `gorm:"type:text" json:"description"`
    ParentID    *uint      `json:"parent_id"`
    IsActive    bool       `gorm:"default:true" json:"is_active"`
    CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
    UpdatedAt   time.Time  `gorm:"default:now()" json:"updated_at"`
    IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
    DeletedAt   *time.Time `json:"deleted_at"`
    DeletedBy   *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

    CodDepSGD *string `gorm:"type:varchar(5)" json:"cod_dep_sgd"`

    // FK necesaria
    Parent *OrganicUnit `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent,omitempty"`
}