import { DeviceInfo } from '@/types';

interface IpWhoResponse {
  success: boolean;
  ip: string;
  country: string;
  country_code: string;
  region: string;
  city: string;
  latitude: number;
  longitude: number;
  connection: {
    isp: string;
    org: string;
  };
  timezone: {
    id: string;
  };
}

interface HttpbunResponse {
  origin: string;
  headers: Record<string, string>;
}

// Helper to parse userAgent for browser and OS info
function parseUA(ua: string) {
  // Browser
  const browserMatch = ua.match(/(Chrome|Firefox|Safari|Edge)\/([\d\.]+)/i);
  const browser_name = browserMatch ? browserMatch[1] : 'Unknown';
  const browser_version = browserMatch ? browserMatch[2] : '0';

  // OS
  const osMatch = ua.match(/\(([^)]+)\)/);
  let os = 'Unknown';
  let os_version = '0';
  if (osMatch) {
    const parts = osMatch[1].split(';')[0].split(' ');
    os = parts[0] || os;
    os_version = parts[1] || '0';
  }

  return { browser_name, browser_version, os, os_version };
}

export async function getDeviceInfo(): Promise<DeviceInfo> {
  // Fetch IP and location data
  const ipRes = await fetch('https://ipwho.is/');
  const ipJson = (await ipRes.json()) as IpWhoResponse;

  // Fetch headers including User-Agent and Accept-Language
  const httpRes = await fetch('https://httpbun.com/get');
  const httpJson = (await httpRes.json()) as HttpbunResponse;

  const ua = httpJson.headers['user-agent'] || httpJson.headers['User-Agent'] || navigator.userAgent;
  const language = httpJson.headers['accept-language'] || navigator.language;
  const { browser_name, browser_version, os, os_version } = parseUA(ua);
  const device_type = /Mobi|Android|iPhone/i.test(ua) ? 'mobile' : 'desktop';
  const device_id = typeof crypto !== 'undefined' && crypto.randomUUID ? crypto.randomUUID() : Math.random().toString(36).substring(2);

  return {
    user_agent: ua,
    ip: ipJson.ip || httpJson.origin,
    device_id,
    browser_name,
    browser_version,
    os,
    os_version,
    device_type,
    timezone: ipJson.timezone.id,
    language,
    location: {
      country: ipJson.country,
      country_code: ipJson.country_code,
      region: ipJson.region,
      city: ipJson.city,
      coordinates: [ipJson.latitude, ipJson.longitude],
      isp: ipJson.connection.isp,
      organization: ipJson.connection.org,
    },
  };
}
