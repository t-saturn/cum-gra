package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrganicUnit struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string
	Acronym     string
	Brand       string
	Level       string
	Description string
	ParentID    *uuid.UUID
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
	DeletedAt   *time.Time
	DeletedBy   *uuid.UUID
}
