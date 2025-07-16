package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type CreateApplicationDTO struct {
	Name         string         `json:"name" validate:"required,min=3,max=100"`
	ClientID     string         `json:"client_id" validate:"required,min=3,max=255"`
	ClientSecret string         `json:"client_secret" validate:"required,min=8,max=255"`
	Domain       string         `json:"domain" validate:"required,url"`
	Logo         *string        `json:"logo" validate:"omitempty,url"`
	Description  *string        `json:"description" validate:"omitempty,max=1000"`
	CallbackUrls pq.StringArray `json:"callback_urls" validate:"required,dive,required,url"`
	IsFirstParty *bool          `json:"is_first_party" validate:"omitempty"`
	Status       *string        `json:"status" validate:"omitempty,oneof=active inactive suspended blocked"`
}

type ApplicationResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	ClientID     string    `json:"client_id"`
	Domain       string    `json:"domain"`
	Logo         *string   `json:"logo,omitempty"`
	Description  *string   `json:"description,omitempty"`
	CallbackUrls []string  `json:"callback_urls"`
	IsFirstParty bool      `json:"is_first_party"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateApplicationDTO struct {
	Name         *string        `json:"name" validate:"omitempty,min=3,max=100"`
	ClientID     *string        `json:"client_id" validate:"omitempty,min=3,max=255"`
	ClientSecret *string        `json:"client_secret" validate:"omitempty,min=8,max=255"`
	Domain       *string        `json:"domain" validate:"omitempty,url"`
	Logo         *string        `json:"logo" validate:"omitempty,url"`
	Description  *string        `json:"description" validate:"omitempty,max=1000"`
	CallbackUrls pq.StringArray `json:"callback_urls" validate:"omitempty,dive,required,url"`
	IsFirstParty *bool          `json:"is_first_party" validate:"omitempty"`
	Status       *string        `json:"status" validate:"omitempty,oneof=active inactive suspended blocked"`
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type CreateApplicationRoleDTO struct {
	Name          string    `json:"name" validate:"required,min=3,max=100"`
	Description   *string   `json:"description" validate:"omitempty,max=1000"`
	ApplicationID uuid.UUID `json:"application_id" validate:"required"`
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type CreateUserApplicationRoleDTO struct {
	UserID            uuid.UUID `json:"user_id" validate:"required"`
	ApplicationID     uuid.UUID `json:"application_id" validate:"required"`
	ApplicationRoleID uuid.UUID `json:"application_role_id" validate:"required"`
	GrantedBy         uuid.UUID `json:"granted_by" validate:"required"`
}
