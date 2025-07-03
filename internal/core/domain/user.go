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

type StructuralPosition struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Code        string     `gorm:"type:varchar(50);unique;not null" json:"code"`
	Level       *int       `gorm:"type:integer; not null" json:"level"`
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
