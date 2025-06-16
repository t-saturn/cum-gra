package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type OAuthToken struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null"`
	ApplicationID uuid.UUID      `gorm:"type:uuid;not null"`
	AccessToken   string         `gorm:"type:varchar(512);unique;not null"`
	RefreshToken  string         `gorm:"type:varchar(512);unique"`
	TokenType     string         `gorm:"type:varchar(50);default:'Bearer'"`
	Scopes        pq.StringArray `gorm:"type:text[]"`
	ExpiresAt     time.Time      `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	RevokedAt     *time.Time

	// Relaciones
	User        User        `gorm:"foreignKey:UserID"`
	Application Application `gorm:"foreignKey:ApplicationID"`
}

func (ot *OAuthToken) BeforeCreate(tx *gorm.DB) error {
	if ot.ID == uuid.Nil {
		ot.ID = uuid.New()
	}
	return nil
}
