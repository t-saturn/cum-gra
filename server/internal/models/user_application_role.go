package models

import (
	"time"

	"github.com/google/uuid"
)

type UserApplicationRole struct {
    ID                uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
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
    CreatedAt         time.Time  `gorm:"default:now()" json:"created_at"`
    UpdatedAt         time.Time  `gorm:"default:now()" json:"updated_at"`

    User            *User            `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
    Application     *Application     `gorm:"foreignKey:ApplicationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"application,omitempty"`
    ApplicationRole *ApplicationRole `gorm:"foreignKey:ApplicationRoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"application_role,omitempty"`
}
