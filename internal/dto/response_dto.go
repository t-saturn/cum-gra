// Package dto define estructuras de datos utilizadas para la entrada y salida en las operaciones de la API.
package dto

type CoordinatesDTO [2]float64

type LocationDetailDTO struct {
	Country      string         `json:"country,omitempty"`
	CountryCode  string         `json:"country_code,omitempty"`
	Region       string         `json:"region,omitempty"`
	City         string         `json:"city,omitempty"`
	Coordinates  CoordinatesDTO `json:"coordinates,omitempty"`
	ISP          string         `json:"isp,omitempty"`
	Organization string         `json:"organization,omitempty"`
}

type DeviceInfoDTO struct {
	UserAgent      string             `json:"user_agent,omitempty"`
	IP             string             `json:"ip,omitempty"`
	DeviceID       string             `json:"device_id,omitempty"`
	BrowserName    string             `json:"browser_name,omitempty"`
	BrowserVersion string             `json:"browser_version,omitempty"`
	OS             string             `json:"os,omitempty"`
	OSVersion      string             `json:"os_version,omitempty"`
	DeviceType     string             `json:"device_type,omitempty"`
	Timezone       string             `json:"timezone,omitempty"`
	Language       string             `json:"language,omitempty"`
	Location       *LocationDetailDTO `json:"location,omitempty"`
}

// ValidationErrorResponse representa una respuesta con errores de validación por campo.
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// HealthResponse representa la respuesta del endpoint /health
type HealthResponse struct {
	Status  string `json:"status"`  // e.g. "ok"
	Message string `json:"message"` // texto adicional opcional
}
