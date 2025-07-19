package models

import (
	"time"

	"github.com/google/uuid"
)

// PasswordHistory guarda un historial de contraseñas anteriores asociadas a un usuario.
type PasswordHistory struct {
	ID                   uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID               uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	PreviousPasswordHash string     `gorm:"type:varchar(255);not null" json:"-"`
	ChangedAt            time.Time  `gorm:"default:now()" json:"changed_at"`
	ChangedBy            *uuid.UUID `gorm:"type:uuid" json:"changed_by"`
	IsDeleted            bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt            *time.Time `json:"deleted_at"`
	DeletedBy            *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	User          *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ChangedByUser *User `gorm:"foreignKey:ChangedBy" json:"changed_by_user,omitempty"`
	DeletedByUser *User `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
}
