package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Email                string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash         string     `gorm:"type:varchar(255);not null" json:"-"`
	FirstName            *string    `gorm:"type:varchar(100)" json:"first_name"`
	LastName             *string    `gorm:"type:varchar(100)" json:"last_name"`
	Phone                *string    `gorm:"type:varchar(20)" json:"phone"`
	DNI                  string     `gorm:"type:varchar(8);unique;not null" json:"dni"`
	EmailVerified        bool       `gorm:"default:false" json:"email_verified"`
	PhoneVerified        bool       `gorm:"default:false" json:"phone_verified"`
	TwoFactorEnabled     bool       `gorm:"default:false" json:"two_factor_enabled"`
	Status               string     `gorm:"type:status_enum;default:'active'" json:"status"`
	StructuralPositionID *uuid.UUID `gorm:"type:uuid" json:"structural_position_id"`
	OrganicUnitID        *uuid.UUID `gorm:"type:uuid" json:"organic_unit_id"`
	CreatedAt            time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"default:now()" json:"updated_at"`
	IsDeleted            bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt            *time.Time `json:"deleted_at"`
	DeletedBy            *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	StructuralPosition     *StructuralPosition     `gorm:"foreignKey:StructuralPositionID" json:"structural_position,omitempty"`
	OrganicUnit            *OrganicUnit            `gorm:"foreignKey:OrganicUnitID" json:"organic_unit,omitempty"`
	DeletedByUser          *User                   `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	UserApplicationRoles   []UserApplicationRole   `gorm:"foreignKey:UserID" json:"user_application_roles,omitempty"`
	UserModuleRestrictions []UserModuleRestriction `gorm:"foreignKey:UserID" json:"user_module_restrictions,omitempty"`
	PasswordHistory        []PasswordHistory       `gorm:"foreignKey:UserID" json:"password_histories,omitempty"`
}
