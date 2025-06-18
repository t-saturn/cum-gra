package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

type User struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Email            string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash     string     `gorm:"type:varchar(255);not null;column:password_hash" json:"-"`
	FirstName        string     `gorm:"type:varchar(100)" json:"first_name"`
	LastName         string     `gorm:"type:varchar(100)" json:"last_name"`
	Phone            string     `gorm:"type:varchar(20)" json:"phone"`
	DNI              string     `gorm:"type:varchar(8);uniqueIndex;not null" json:"dni"`
	EmailVerified    bool       `gorm:"default:false" json:"email_verified"`
	PhoneVerified    bool       `gorm:"default:false" json:"phone_verified"`
	TwoFactorEnabled bool       `gorm:"default:false" json:"two_factor_enabled"`
	Status           UserStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	/* Status           UserStatus `gorm:"type:enum('active','suspended','deleted');default:'active'" json:"status"` */

	StructuralPositionID uuid.UUID           `gorm:"type:uuid" json:"structural_position_id"`
	StructuralPosition   *StructuralPosition `gorm:"foreignKey:StructuralPositionID" json:"-"`
	OrganicUnitID        uuid.UUID           `gorm:"type:uuid" json:"organic_unit_id"`
	OrganicUnit          *OrganicUnit        `gorm:"foreignKey:OrganicUnitID" json:"-"`

	LastLoginAt *time.Time `json:"last_login_at"`

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	IsDeleted bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`
}
