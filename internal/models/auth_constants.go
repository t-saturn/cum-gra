package models

const (
	AuthMethodCredentials = "credentials"
	AuthMethodToken       = "token"

	AuthStatusPending = "pending"
	AuthStatusSuccess = "success"
	AuthStatusFailed  = "failed"
	AuthStatusExpired = "expired"

	TokenStatusActive  = "active"
	TokenStatusRevoked = "revoked"
	TokenStatusExpired = "expired"
	TokenStatusInvalid = "invalid"
)
