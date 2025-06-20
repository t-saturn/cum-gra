package dto

import "time"

type CreatePasswordHistoryDTO struct {
	UserID               string    `json:"user_id" validate:"required,uuid"`
	PreviousPasswordHash string    `json:"previous_password_hash" validate:"required"`
	ChangedAt            time.Time `json:"changed_at"`
	ChangedBy            string    `json:"changed_by" validate:"required,uuid"`
}

type UpdatePasswordHistoryDTO struct {
	PreviousPasswordHash string    `json:"previous_password_hash" validate:"required"`
	ChangedAt            time.Time `json:"changed_at"`
	ChangedBy            string    `json:"changed_by" validate:"required,uuid"`
}
