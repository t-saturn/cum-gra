package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	ApplicationActive    ApplicationStatus = "active"
	ApplicationSuspended ApplicationStatus = "suspended"
)

type Application struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey"`
	Name         string            `gorm:"type:varchar(100);not null"`
	ClientID     string            `gorm:"type:varchar(255);unique;not null"`
	ClientSecret string            `gorm:"type:varchar(255);not null"`
	Domain       string            `gorm:"type:varchar(255);not null"`
	Logo         string            `gorm:"type:varchar(255)"`
	Description  string            `gorm:"type:text"`
	CallbackURLs pq.StringArray    `gorm:"type:text[]"`
	Scopes       pq.StringArray    `gorm:"type:text[]"`
	IsFirstParty bool              `gorm:"default:false"`
	Status       ApplicationStatus `gorm:"type:varchar(50);default:'active'"`
	CreatedAt    time.Time         `gorm:"autoCreateTime"`
	UpdatedAt    time.Time         `gorm:"autoUpdateTime"`
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
