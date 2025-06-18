package domain

import (
	"time"

	"github.com/google/uuid"
)

type StructuralPosition struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Code        string    `gorm:"type:varchar(50);uniqueIndex" validate:"required"`
	Level       string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
	DeletedAt   *time.Time
	DeletedBy   *uuid.UUID
}
