package models

// Coordenadas geográficas
type Coordinates [2]float64

// Ubicación del dispositivo
type Location struct {
	Country     string      `bson:"country,omitempty"`
	City        string      `bson:"city,omitempty"`
	Coordinates Coordinates `bson:"coordinates,omitempty"`
}

// Información del dispositivoauth_attempt.go
type DeviceInfo struct {
	UserAgent      string    `bson:"userAgent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"deviceId,omitempty"`
	BrowserName    string    `bson:"browserName,omitempty"`
	BrowserVersion string    `bson:"browserVersion,omitempty"`
	OS             string    `bson:"os,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}
