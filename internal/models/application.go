// Package models contiene las definiciones de los modelos utilizados por GORM para mapear la base de datos.
package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Application representa una aplicación o sistema registrado en la plataforma.
type Application struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name"`
	ClientID     string         `gorm:"type:varchar(255);unique;not null" json:"client_id"`
	ClientSecret string         `gorm:"type:varchar(255);not null" json:"-"`
	Domain       string         `gorm:"type:varchar(255);not null" json:"domain"`
	Logo         *string        `gorm:"type:varchar(255)" json:"logo"`
	Description  *string        `gorm:"type:text" json:"description"`
	CallbackUrls pq.StringArray `gorm:"type:text[]" json:"callback_urls"`
	Status       string         `gorm:"type:application_status_enum;default:'active'" json:"status"`
	CreatedAt    time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"default:now()" json:"updated_at"`
	IsDeleted    bool           `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt    *time.Time     `json:"deleted_at"`
	DeletedBy    *uuid.UUID     `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	DeletedByUser          *User                   `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	ApplicationRoles       []ApplicationRole       `gorm:"foreignKey:ApplicationID" json:"application_roles,omitempty"`
	UserApplicationRoles   []UserApplicationRole   `gorm:"foreignKey:ApplicationID" json:"user_application_roles,omitempty"`
	Modules                []Module                `gorm:"foreignKey:ApplicationID" json:"modules,omitempty"`
	UserModuleRestrictions []UserModuleRestriction `gorm:"foreignKey:ApplicationID" json:"user_module_restrictions,omitempty"`
}
