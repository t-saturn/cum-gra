package models

// Coordenadas geográficas
type Coordinates [2]float64

// Ubicación del dispositivo
type Location struct {
	Country     string      `bson:"country,omitempty"`
	City        string      `bson:"city,omitempty"`
	Coordinates Coordinates `bson:"coordinates,omitempty"`
}

// Información del dispositivo
type DeviceInfo struct {
	UserAgent      string    `bson:"user_agent,omitempty"`
	IP             string    `bson:"ip,omitempty"`
	DeviceID       string    `bson:"device_id,omitempty"`
	BrowserName    string    `bson:"browser_name,omitempty"`
	BrowserVersion string    `bson:"browser_version,omitempty"`
	OS             string    `bson:"os,omitempty"`
	Location       *Location `bson:"location,omitempty"`
}
