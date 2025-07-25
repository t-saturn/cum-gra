package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- ENUMS ---

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

// --- COMUNES ---
type Coordinates [2]float64

type Location struct {
	Country     string      `bson:"country,omitempty"`
	City        string      `bson:"city,omitempty"`
	Coordinates Coordinates `bson:"coordinates,omitempty"`
}

type DeviceInfo struct {
	UserAgent      string    `bson:"user_agent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"device_id,omitempty"`
	BrowserName    string    `bson:"browser_name,omitempty"`
	BrowserVersion string    `bson:"browser_version,omitempty"`
	OS             string    `bson:"os,omitempty"`
	OSVersion      string    `bson:"os_version,omitempty"`
	DeviceType     string    `bson:"device_type,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}

// --- CAPTCHA LOG ---
type CaptchaLog struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	Token         string              `bson:"token"`
	Success       bool                `bson:"success"`
	ChallengeTS   time.Time           `bson:"challenge_ts"`
	Hostname      string              `bson:"hostname"`
	Action        string              `bson:"action,omitempty"`
	CustomData    string              `bson:"cdata,omitempty"`
	ErrorCodes    []string            `bson:"error_codes,omitempty"`
	RemoteIP      string              `bson:"remote_ip"`
	CreatedAt     time.Time           `bson:"created_at"`
	UserID        primitive.ObjectID  `bson:"user_id,omitempty"`
	SessionID     *primitive.ObjectID `bson:"session_id,omitempty"`
	AuthLogID     *primitive.ObjectID `bson:"auth_log_id,omitempty"`
	ApplicationID string              `bson:"application_id,omitempty"`
}
