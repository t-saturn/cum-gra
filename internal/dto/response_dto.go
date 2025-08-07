// Package dto define estructuras de datos utilizadas para la entrada y salida en las operaciones de la API.
package dto

import "github.com/t-saturn/auth-service-server/internal/models"

// ErrorDTO describe la información de error incluida en las respuestas de la API.
type ErrorDTO struct {
	Code    string `json:"code"`
	Details string `json:"details,omitempty"`
}

type ResponseDTO[T any] struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    T         `json:"data"`
	Error   *ErrorDTO `json:"error"`
}

type ValidationResponseDTO struct {
	UserID          string `bson:"user_id,omitempty"`
	ServiceResponse string `bson:"service_response,omitempty"`
	ValidatedBy     string `bson:"validated_by,omitempty"`
	ValidationTime  int64  `bson:"validation_time,omitempty"` // tiempo en ms
}

// ValidationErrorResponse representa una respuesta con errores de validación por campo.
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

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

// ToModel convierte un DeviceInfoDTO a models.DeviceInfo
func (d DeviceInfoDTO) ToModel() models.DeviceInfo {
	mi := models.DeviceInfo{
		UserAgent:      d.UserAgent,
		IP:             d.IP,
		DeviceID:       d.DeviceID,
		BrowserName:    d.BrowserName,
		BrowserVersion: d.BrowserVersion,
		OS:             d.OS,
		OSVersion:      d.OSVersion,
		DeviceType:     d.DeviceType,
		Timezone:       d.Timezone,
		Language:       d.Language,
	}
	if d.Location != nil {
		mi.Location = &models.LocationDetail{
			Country:      d.Location.Country,
			CountryCode:  d.Location.CountryCode,
			Region:       d.Location.Region,
			City:         d.Location.City,
			Coordinates:  models.Coordinates{d.Location.Coordinates[0], d.Location.Coordinates[1]},
			ISP:          d.Location.ISP,
			Organization: d.Location.Organization,
		}
	}
	return mi
}
