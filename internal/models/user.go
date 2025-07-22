package models

import (
	"time"

	"github.com/google/uuid"
)

// User representa a un usuario del sistema, incluyendo su perfil, credenciales y relaciones.
type User struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Email            string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash     string     `gorm:"type:varchar(255);not null" json:"-"`
	DNI              string     `gorm:"type:varchar(8);unique;not null" json:"dni"`
	EmailVerified    bool       `gorm:"default:false" json:"email_verified"`
	TwoFactorEnabled bool       `gorm:"default:false" json:"two_factor_enabled"`
	Status           string     `gorm:"type:status_enum;default:'active'" json:"status"`
	IsDeleted        bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt        *time.Time `json:"deleted_at"`
	DeletedBy        *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`
}
