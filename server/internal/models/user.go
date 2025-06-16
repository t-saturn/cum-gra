package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStatus string

const (
	UserActive    UserStatus = "active"
	UserSuspended UserStatus = "suspended"
	UserDeleted   UserStatus = "deleted"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     *string   `gorm:"type:varchar(255);unique"`
	Password  *string   `gorm:"type:varchar(255);column:password_hash"`
	FirstName string    `gorm:"type:varchar(100)"`
	LastName  string    `gorm:"type:varchar(100)"`
	Phone     string    `gorm:"type:varchar(20)"`

	// Verificaciones
	EmailVerified bool `gorm:"default:false"`
	PhoneVerified bool `gorm:"default:false"`

	// Configuración de seguridad
	TwoFactorEnabled      bool       `gorm:"default:false"`
	RequireBiometric      bool       `gorm:"default:false"`
	PreferredAuthMethodID *uuid.UUID `gorm:"type:uuid"`

	// Status y estructura organizacional
	Status               UserStatus `gorm:"type:varchar(50);default:'active'"`
	StructuralPositionID *uuid.UUID `gorm:"type:uuid"`
	OrganicUnitID        *uuid.UUID `gorm:"type:uuid"`

	// Timestamps
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	LastLoginAt *time.Time

	// Borrado lógico
	IsDeleted bool `gorm:"not null;default:false"`
	DeletedAt *time.Time
	DeletedBy *uuid.UUID `gorm:"type:uuid"`

	// Relaciones
	//PreferredAuthMethod *AuthenticationMethod `gorm:"foreignKey:PreferredAuthMethodID"`
	StructuralPosition *StructuralPosition `gorm:"foreignKey:StructuralPositionID"`
	OrganicUnit        *OrganicUnit        `gorm:"foreignKey:OrganicUnitID"`
	DeletedByUser      *User               `gorm:"foreignKey:DeletedBy"`

	// Relaciones inversas
	UserSessions            []UserSession         `gorm:"foreignKey:UserID"`
	UserApplicationRoles    []UserApplicationRole `gorm:"foreignKey:UserID"`
	OAuthTokens             []OAuthToken          `gorm:"foreignKey:UserID"`
	UserPermissions         []UserPermission      `gorm:"foreignKey:UserID"`
	PasswordHistory         []PasswordHistory     `gorm:"foreignKey:UserID"`
	PasswordResets          []PasswordReset       `gorm:"foreignKey:UserID"`
	TwoFactorSecrets        []TwoFactorSecret     `gorm:"foreignKey:UserID"`
	AuditLogs               []AuditLog            `gorm:"foreignKey:UserID"`
	UserPreferences         []UserPreference      `gorm:"foreignKey:UserID"`
	GrantedApplicationRoles []UserApplicationRole `gorm:"foreignKey:GrantedBy"`
	RevokedApplicationRoles []UserApplicationRole `gorm:"foreignKey:RevokedBy"`
	GrantedPermissions      []UserPermission      `gorm:"foreignKey:GrantedBy"`
	RevokedPermissions      []UserPermission      `gorm:"foreignKey:RevokedBy"`
	ChangedPasswords        []PasswordHistory     `gorm:"foreignKey:ChangedBy"`
	DeletedPasswordHistory  []PasswordHistory     `gorm:"foreignKey:DeletedBy"`
	DeletedUsers            []User                `gorm:"foreignKey:DeletedBy"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
