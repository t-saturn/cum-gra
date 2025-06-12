package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"type:uuid;primaryKey"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
}

// Hook para asignar UUID automáticamente
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
