package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email  string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	DNI    string    `gorm:"type:varchar(8);unique;not null" json:"dni"`
	Status string    `gorm:"type:varchar(20);default:'active'" json:"status"`

	CreatedAt time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:now()" json:"updated_at"`
	IsDeleted bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `gorm:"type:uuid" json:"deleted_by"`

	// Relaciones
	DeletedByUser          *User                   `gorm:"foreignKey:DeletedBy" json:"deleted_by_user,omitempty"`
	Details                *UserDetail             `gorm:"foreignKey:UserID" json:"details,omitempty"`
	UserApplicationRoles   []UserApplicationRole   `gorm:"foreignKey:UserID" json:"user_application_roles,omitempty"`
	UserModuleRestrictions []UserModuleRestriction `gorm:"foreignKey:UserID" json:"user_module_restrictions,omitempty"`
}

type UserDetail struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	UserID uuid.UUID `gorm:"type:uuid;unique;not null" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`

	FirstName *string `gorm:"type:varchar(100)" json:"first_name"`
	LastName  *string `gorm:"type:varchar(100)" json:"last_name"`
	Phone     *string `gorm:"type:varchar(20)" json:"phone"`

	// Cargo y dependencia (antes en User)
	StructuralPositionID *uint               `json:"structural_position_id"`
	OrganicUnitID        *uint               `json:"organic_unit_id"`
	StructuralPosition   *StructuralPosition `gorm:"foreignKey:StructuralPositionID" json:"structural_position,omitempty"`
	OrganicUnit          *OrganicUnit        `gorm:"foreignKey:OrganicUnitID" json:"organic_unit,omitempty"`

	// Relación con Ubigeo (lugar de residencia, por ejemplo)
	UbigeoID *uint   `json:"ubigeo_id"`
	Ubigeo   *Ubigeo `gorm:"foreignKey:UbigeoID" json:"ubigeo,omitempty"`
}

type Ubigeo struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	UbigeoCode string `gorm:"type:varchar(10);unique;not null" json:"ubigeo_code"`
	IneiCode   string `gorm:"type:varchar(10)" json:"inei_code"`

	Department string `gorm:"type:varchar(100);not null" json:"department"`
	Province   string `gorm:"type:varchar(100);not null" json:"province"`
	District   string `gorm:"type:varchar(100);not null" json:"district"`

	UserDetails []UserDetail `gorm:"foreignKey:UbigeoID" json:"user_details,omitempty"`

	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
}
