package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserApplicationRole struct {
	ID                uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	UserID            uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	ApplicationID     uuid.UUID  `gorm:"type:uuid;not null" json:"application_id"`
	ApplicationRoleID uuid.UUID  `gorm:"type:uuid;not null" json:"application_role_id"`
	GrantedAt         time.Time  `gorm:"default:now()" json:"granted_at"`
	GrantedBy         uuid.UUID  `gorm:"type:uuid;not null" json:"granted_by"`
	RevokedAt         *time.Time `json:"revoked_at"`
	RevokedBy         *uuid.UUID `gorm:"type:uuid" json:"revoked_by"`
	IsDeleted         bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt         *time.Time `json:"deleted_at"`
	DeletedBy         *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	User            *User            `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Application     *Application     `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	ApplicationRole *ApplicationRole `gorm:"foreignKey:ApplicationRoleID" json:"application_role,omitempty"`
	GrantedByUser   *User            `gorm:"foreignKey:GrantedBy" json:"granted_by_user,omitempty"`
	RevokedByUser   *User            `gorm:"foreignKey:RevokedBy" json:"revoked_by_user,omitempty"`
	DeletedByUser   *User            `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
}
