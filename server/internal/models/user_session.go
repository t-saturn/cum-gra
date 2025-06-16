package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSession struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	SessionToken   string    `gorm:"type:varchar(512);unique;not null"`
	RefreshToken   string    `gorm:"type:varchar(512);unique;not null"`
	DeviceInfo     string    `gorm:"type:text"`
	IPAddress      string    `gorm:"type:inet"`
	UserAgent      string    `gorm:"type:text"`
	IsActive       bool      `gorm:"default:true"`
	ExpiresAt      time.Time `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	LastActivityAt time.Time `gorm:"autoCreateTime"`

	// Borrado lógico
	IsDeleted bool `gorm:"not null;default:false"`
	DeletedAt *time.Time

	// Relaciones
	User User `gorm:"foreignKey:UserID"`
}

func (us *UserSession) BeforeCreate(tx *gorm.DB) error {
	if us.ID == uuid.Nil {
		us.ID = uuid.New()
	}
	return nil
}
