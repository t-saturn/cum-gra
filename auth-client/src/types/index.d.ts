export interface DeviceInfo {
  user_agent: string;
  ip: string;
  device_id: string;
  browser_name: string;
  browser_version: string;
  os: string;
  os_version: string;
  device_type: string;
  timezone: string;
  language: string;
  location: Location;
}

export interface Location {
  country: string;
  country_code: string;
  region: string;
  city: string;
  coordinates: number[];
  isp: string;
  organization: string;
}
