package models

func convertLocation(ld *LocationDetail) *Location {
	if ld == nil {
		return nil
	}

	return &Location{
		Country:     ld.Country,
		City:        ld.City,
		Coordinates: ld.Coordinates,
	}
}

func (s SessionDeviceInfo) ToDeviceInfo() DeviceInfo {
	return DeviceInfo{
		UserAgent:      s.UserAgent,
		IP:             s.IP,
		DeviceID:       s.DeviceID,
		BrowserName:    s.BrowserName,
		BrowserVersion: s.BrowserVersion,
		OS:             s.OS,
		Location:       convertLocation(s.Location),
	}
}
