package dto

import (
	"time"
)

// ─────────────────────────────────────────────────────────────────────────────
// Vistas / DTOs base
// ─────────────────────────────────────────────────────────────────────────────

type SessionViewDTO struct {
	SessionID       string        `json:"session_id"`
	UserID          string        `json:"user_id"`
	Status          string        `json:"status"` // active | inactive | revoked | expired
	IsActive        bool          `json:"is_active"`
	MaxRefreshAt    time.Time     `json:"max_refresh_at"`
	CreatedAt       time.Time     `json:"created_at"`
	LastActivity    time.Time     `json:"last_activity"`
	ExpiresAt       time.Time     `json:"expires_at"`
	RevokedAt       *time.Time    `json:"revoked_at,omitempty"`
	DeviceInfo      DeviceInfoDTO `json:"device_info"`
	TokensGenerated []string      `json:"tokens_generated,omitempty"` // ObjectIDs en hex
}

type PaginationMeta struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	HasPrev bool  `json:"has_prev"`
	HasNext bool  `json:"has_next"`
}

// ─────────────────────────────────────────────────────────────────────────────
// GET /auth/me?session_id=<session_id>
// ─────────────────────────────────────────────────────────────────────────────

type AuthMeQueryDTO struct {
	SessionID string `query:"session_id" json:"session_id" validate:"required"` // ID de la sesión asociada
}

type AuthMeResponseDTO struct {
	UserID             string         `json:"user_id"`
	Session            SessionViewDTO `json:"session"`
	Email              string         `json:"email,omitempty"`
	Name               string         `json:"name,omitempty"`
	DNI                string         `json:"dni,omitempty"`   // <- NUEVO (o *string si quieres null)
	Phone              string         `json:"phone,omitempty"` // <- (o *string si quieres null)
	Status             string         `json:"status,omitempty"`
	StructuralPosition string         `json:"structural_position,omitempty"`
	OrganicUnit        string         `json:"organic_unit,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// GET /auth/sessions?session_id=<session_id>&page=<page>&limit=<limit>&sort_by=<sort_by>&sort_order=<sort_order> (del usuario autenticado)
// ─────────────────────────────────────────────────────────────────────────────

type ListSessionsQueryDTO struct {
	SessionID   string         `query:"session_id" json:"session_id" validate:"required"` // ID de la sesión asociada
	Status      *string        `query:"status" json:"status,omitempty"`                   // active|inactive|revoked|expired
	IsActive    *bool          `query:"is_active" json:"is_active,omitempty"`
	CreatedFrom *time.Time     `query:"created_from" json:"created_from,omitempty"` // RFC3339
	CreatedTo   *time.Time     `query:"created_to"   json:"created_to,omitempty"`
	Session     SessionViewDTO `json:"session"`
	Page        int            `query:"page" json:"page"`             // default 1
	Limit       int            `query:"limit" json:"limit"`           // default 20
	SortBy      string         `query:"sort_by" json:"sort_by"`       // default "created_at"
	SortOrder   string         `query:"sort_order" json:"sort_order"` // asc|desc, default "desc"
}

type ListSessionsResponseDTO struct {
	Data       []SessionViewDTO `json:"data"`
	Pagination PaginationMeta   `json:"pagination"`
}

// ─────────────────────────────────────────────────────────────────────────────
// DELETE /auth/sessions?session_id=<session_id>&reason=<reason>&revoked_by_app=<revoked_by_app>
// ─────────────────────────────────────────────────────────────────────────────

type RevokeOwnSessionQueryDTO struct {
	SessionID    string  `query:"session_id" json:"session_id" validate:"required"` // ID de la sesión asociada
	Reason       *string `query:"reason" json:"reason,omitempty" validate:"required,oneof=user_logout refresh_token session_expired admin_revoked forbidden_revoke"`
	RevokedByApp *string `query:"revoked_by_app" json:"revoked_by_app,omitempty"`
}

type RevokeOwnSessionResponseDTO struct {
	SessionID string                  `json:"session_id"`
	Status    string                  `json:"status"` // "revoked" esperado
	RevokedAt *time.Time              `json:"revoked_at,omitempty"`
	Message   string                  `json:"message"`
	Tokens    []RevokedTokenDetailDTO `json:"revoked_tokens,omitempty"`
}

type RevokedTokenDetailDTO struct {
	TokenID   string     `json:"token_id"`
	TokenType string     `json:"token_type"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	Reason    string     `json:"reason,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// GET /admin/sessions
// ─────────────────────────────────────────────────────────────────────────────

type AdminListSessionsQueryDTO struct {
	UserID           *string    `query:"user_id" json:"user_id,omitempty"`
	SessionID        string     `query:"session_id" json:"session_id" validate:"required"`
	Status           *string    `query:"status" json:"status,omitempty"` // active|inactive|revoked|expired
	IsActive         *bool      `query:"is_active" json:"is_active,omitempty"`
	CreatedFrom      *time.Time `query:"created_from" json:"created_from,omitempty"`
	CreatedTo        *time.Time `query:"created_to" json:"created_to,omitempty"`
	LastActivityFrom *time.Time `query:"last_activity_from" json:"last_activity_from,omitempty"`
	LastActivityTo   *time.Time `query:"last_activity_to" json:"last_activity_to,omitempty"`
	Page             int        `query:"page" json:"page"`             // default 1
	Limit            int        `query:"limit" json:"limit"`           // default 20
	SortBy           string     `query:"sort_by" json:"sort_by"`       // default "created_at"
	SortOrder        string     `query:"sort_order" json:"sort_order"` // asc|desc, default "desc"
}

type AdminListSessionsResponseDTO struct {
	Data       []SessionViewDTO `json:"data"`
	Pagination PaginationMeta   `json:"pagination"`
}

// ─────────────────────────────────────────────────────────────────────────────
// DELETE /admin/sessions?session_id=
// ─────────────────────────────────────────────────────────────────────────────

type AdminRevokeSessionQueryDTO struct {
	SessionID    string  `query:"session_id" json:"session_id" validate:"required"`
	Reason       *string `query:"reason" json:"reason,omitempty"`
	RevokedBy    *string `query:"revoked_by" json:"revoked_by,omitempty"`         // ObjectID de admin o user que revoca
	RevokedByApp *string `query:"revoked_by_app" json:"revoked_by_app,omitempty"` // app/servicio que revoca
}

type AdminRevokeSessionResponseDTO struct {
	SessionID string     `json:"session_id"`
	Status    string     `json:"status"` // "revoked"
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	Message   string     `json:"message"`
}

// ─────────────────────────────────────────────────────────────────────────────
// DELETE /admin/sessions?user_id=   (revocar en bloque del usuario)
// ─────────────────────────────────────────────────────────────────────────────

type AdminRevokeUserSessionsQueryDTO struct {
	UserID           string   `query:"user_id" json:"user_id" validate:"required"`
	SessionID        string   `query:"session_id" json:"session_id" validate:"required"`
	Reason           *string  `query:"reason" json:"reason,omitempty"`
	RevokedBy        *string  `query:"revoked_by" json:"revoked_by,omitempty"`
	RevokedByApp     *string  `query:"revoked_by_app" json:"revoked_by_app,omitempty"`
	IncludeStatuses  []string `query:"include_statuses" json:"include_statuses,omitempty"`     // si quieres filtrar qué estados revocar
	ExcludeSessionID *string  `query:"exclude_session_id" json:"exclude_session_id,omitempty"` // p.ej. no tumbar la actual
}

type AdminBulkRevokeResponseDTO struct {
	UserID        string   `json:"user_id"`
	AffectedCount int      `json:"affected_count"`
	Revoked       []string `json:"revoked_session_ids"`
	Message       string   `json:"message"`
}
