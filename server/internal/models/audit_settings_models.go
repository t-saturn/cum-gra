package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuditLog model
type AuditLog struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null"`
	ApplicationID uuid.UUID  `gorm:"type:uuid;not null"`
	Action        string     `gorm:"type:varchar(100);not null"`
	ResourceType  *string    `gorm:"type:varchar(50)"`
	ResourceID    *uuid.UUID `gorm:"type:uuid"`
	IPAddress     string     `gorm:"type:inet"`
	UserAgent     string     `gorm:"type:text"`
	//Details       JSONB      `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relaciones
	User        User        `gorm:"foreignKey:UserID"`
	Application Application `gorm:"foreignKey:ApplicationID"`
}

// ApplicationSetting model
type SettingDataType string

const (
	SettingString  SettingDataType = "string"
	SettingNumber  SettingDataType = "number"
	SettingBoolean SettingDataType = "boolean"
	SettingJSON    SettingDataType = "json"
)

type ApplicationSetting struct {
	ID            uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ApplicationID uuid.UUID       `gorm:"type:uuid;not null"`
	SettingKey    string          `gorm:"type:varchar(100);not null"`
	SettingValue  string          `gorm:"type:text"`
	DataType      SettingDataType `gorm:"type:varchar(50);default:'string'"`
	Description   string          `gorm:"type:text"`
	IsPublic      bool            `gorm:"default:false"`
	CreatedAt     time.Time       `gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `gorm:"autoUpdateTime"`

	// Relaciones
	Application Application `gorm:"foreignKey:ApplicationID"`
}

// UserPreference model
type UserPreference struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	ApplicationID   uuid.UUID `gorm:"type:uuid;not null"`
	PreferenceKey   string    `gorm:"type:varchar(100);not null"`
	PreferenceValue string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	// Relaciones
	User        User        `gorm:"foreignKey:UserID"`
	Application Application `gorm:"foreignKey:ApplicationID"`
}

// BeforeCreate hooks
func (al *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if al.ID == uuid.Nil {
		al.ID = uuid.New()
	}
	return nil
}

func (as *ApplicationSetting) BeforeCreate(tx *gorm.DB) error {
	if as.ID == uuid.Nil {
		as.ID = uuid.New()
	}
	return nil
}

func (up *UserPreference) BeforeCreate(tx *gorm.DB) error {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return nil
}
