package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserStatus representa los estados posibles del usuario
type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

// User representa el modelo de usuario en la base de datos
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null;column:password_hash" json:"-"`
	FirstName    *string   `gorm:"type:varchar(100)" json:"first_name,omitempty"`
	LastName     *string   `gorm:"type:varchar(100)" json:"last_name,omitempty"`
	Phone        *string   `gorm:"type:varchar(20)" json:"phone,omitempty"`

	// Verificaciones
	EmailVerified    bool `gorm:"default:false" json:"email_verified"`
	PhoneVerified    bool `gorm:"default:false" json:"phone_verified"`
	TwoFactorEnabled bool `gorm:"default:false" json:"two_factor_enabled"`

	// Status y estructura organizacional
	Status               UserStatus `gorm:"type:varchar(50);default:'active'" json:"status"`
	StructuralPositionID *uuid.UUID `gorm:"type:uuid" json:"structural_position_id,omitempty"`
	OrganicUnitID        *uuid.UUID `gorm:"type:uuid" json:"organic_unit_id,omitempty"`

	// Timestamps
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`

	// Borrado lógico
	IsDeleted bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"type:uuid" json:"deleted_by,omitempty"`

	// Relaciones
	DeletedByUser *User `gorm:"foreignKey:DeletedBy;references:ID" json:"deleted_by_user,omitempty"`

	// Relaciones inversas (usuarios eliminados por este usuario)
	DeletedUsers []User `gorm:"foreignKey:DeletedBy" json:"deleted_users,omitempty"`
}

// BeforeCreate hook de GORM para generar UUID antes de crear
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// TableName especifica el nombre de la tabla
func (User) TableName() string {
	return "users"
}

// IsActive verifica si el usuario está activo
func (u *User) IsActive() bool {
	return u.Status == UserStatusActive && !u.IsDeleted
}

// GetFullName retorna el nombre completo del usuario
func (u *User) GetFullName() string {
	var fullName string
	if u.FirstName != nil {
		fullName = *u.FirstName
	}
	if u.LastName != nil {
		if fullName != "" {
			fullName += " "
		}
		fullName += *u.LastName
	}
	return fullName
}

// SoftDelete realiza un borrado lógico del usuario
func (u *User) SoftDelete(deletedBy uuid.UUID) {
	now := time.Now()
	u.IsDeleted = true
	u.DeletedAt = &now
	u.DeletedBy = &deletedBy
	u.Status = UserStatusDeleted
}

// Restore restaura un usuario eliminado lógicamente
func (u *User) Restore() {
	u.IsDeleted = false
	u.DeletedAt = nil
	u.DeletedBy = nil
	u.Status = UserStatusActive
}

// UpdateLastLogin actualiza el timestamp del último login
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
}
