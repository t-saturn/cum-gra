package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModuleRestriction struct {
    ID                 uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
    UserID             uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
    ModuleID           uuid.UUID  `gorm:"type:uuid;not null" json:"module_id"`
    ApplicationID      uuid.UUID  `gorm:"type:uuid;not null" json:"application_id"`
    RestrictionType    string     `gorm:"type:varchar(20);not null" json:"restriction_type"`
    MaxPermissionLevel *string    `gorm:"type:varchar(20)" json:"max_permission_level"`
    Reason             *string    `gorm:"type:text" json:"reason"`
    ExpiresAt          *time.Time `json:"expires_at"`
    CreatedAt          time.Time  `gorm:"default:now()" json:"created_at"`
    CreatedBy          uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
    UpdatedAt          time.Time  `gorm:"default:now()" json:"updated_at"`
    UpdatedBy          *uuid.UUID `gorm:"type:uuid" json:"updated_by"`
    IsDeleted          bool       `gorm:"not null;default:false" json:"is_deleted"`
    DeletedAt          *time.Time `json:"deleted_at"`
    DeletedBy          *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

    User        *User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
    Module      *Module      `gorm:"foreignKey:ModuleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"module,omitempty"`
    Application *Application `gorm:"foreignKey:ApplicationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"application,omitempty"`
}