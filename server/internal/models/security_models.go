package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// PasswordHistory model
type PasswordHistory struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID               uuid.UUID `gorm:"type:uuid;not null"`
	PreviousPasswordHash string    `gorm:"type:varchar(255);not null"`
	ChangedAt            time.Time `gorm:"autoCreateTime"`
	ChangedBy            uuid.UUID `gorm:"type:uuid;not null"`

	// Borrado lógico
	IsDeleted bool `gorm:"not null;default:false"`
	DeletedAt *time.Time
	DeletedBy *uuid.UUID `gorm:"type:uuid"`

	// Relaciones
	User          User  `gorm:"foreignKey:UserID"`
	ChangedByUser User  `gorm:"foreignKey:ChangedBy"`
	DeletedByUser *User `gorm:"foreignKey:DeletedBy"`
}

// PasswordReset model
type PasswordReset struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"type:varchar(255);unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UsedAt    *time.Time
	IPAddress string    `gorm:"type:inet"`
	UserAgent string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relaciones
	User User `gorm:"foreignKey:UserID"`
}

// TwoFactorSecret model
type TwoFactorSecret struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID            uuid.UUID      `gorm:"type:uuid;not null"`
	Secret            string         `gorm:"type:varchar(255);not null"`
	BackupCodes       pq.StringArray `gorm:"type:text[]"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	LastUsedAt        *time.Time
	RecoveryCodesUsed int `gorm:"default:0"`

	// Relaciones
	User User `gorm:"foreignKey:UserID"`
}

// BeforeCreate hooks
func (ph *PasswordHistory) BeforeCreate(tx *gorm.DB) error {
	if ph.ID == uuid.Nil {
		ph.ID = uuid.New()
	}
	return nil
}

func (pr *PasswordReset) BeforeCreate(tx *gorm.DB) error {
	if pr.ID == uuid.Nil {
		pr.ID = uuid.New()
	}
	return nil
}

func (tfs *TwoFactorSecret) BeforeCreate(tx *gorm.DB) error {
	if tfs.ID == uuid.Nil {
		tfs.ID = uuid.New()
	}
	return nil
}
