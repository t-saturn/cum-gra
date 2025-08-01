package dto

import "time"

// HealthResponseDTO define la parte "data" para GET /auth/health
type HealthResponseDTO struct {
	Status       string                       `json:"status"`    // e.g. "healthy"
	Timestamp    time.Time                    `json:"timestamp"` // ISO8601
	Version      string                       `json:"version"`   // e.g. "1.0.0"
	Uptime       string                       `json:"uptime"`    // e.g. "24h15m30s"
	Databases    map[string]DatabaseHealthDTO `json:"databases"`
	Dependencies map[string]string            `json:"dependencies"`
}

// DatabaseHealthDTO describe el estado de una BD
type DatabaseHealthDTO struct {
	Status       string `json:"status"`        // e.g. "connected" | "down"
	ResponseTime string `json:"response_time"` // e.g. "15ms"
}
