import { DeviceInfo } from '@/types';

// Interface que define la respuesta de ipwho.is (información de IP y ubicación)
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

// Interface que define la respuesta de httpbun.com/get (información de headers HTTP)
interface HttpbunResponse {
  origin: string; // IP detectada por httpbun
  headers: Record<string, string>; // Todos los headers enviados en la petición
}

// Helper: parsea el User-Agent para extraer información de navegador y sistema operativo
function parseUA(ua: string) {
  // 1. Detectar navegador y versión
  const browserMatch = ua.match(/(Chrome|Firefox|Safari|Edge)\/([\d\.]+)/i);
  const browser_name = browserMatch ? browserMatch[1] : 'Unknown';
  const browser_version = browserMatch ? browserMatch[2] : '0';

  // 2. Detectar sistema operativo y versión
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

// Función principal: obtiene información del dispositivo y conexión del usuario
export async function getDeviceInfo(): Promise<DeviceInfo> {
  // 1. Obtener datos de IP y ubicación desde ipwho.is
  const ipRes = await fetch('https://ipwho.is/');
  const ipJson = (await ipRes.json()) as IpWhoResponse;

  // 2. Obtener headers (incluyendo User-Agent y Accept-Language) desde httpbun.com
  const httpRes = await fetch('https://httpbun.com/get');
  const httpJson = (await httpRes.json()) as HttpbunResponse;

  // 3. Determinar User-Agent desde httpbun o navegador
  const ua = httpJson.headers['user-agent'] || httpJson.headers['User-Agent'] || navigator.userAgent;

  // 4. Determinar idioma preferido
  const language = httpJson.headers['accept-language'] || navigator.language;

  // 5. Parsear User-Agent para extraer datos de navegador y sistema operativo
  const { browser_name, browser_version, os, os_version } = parseUA(ua);

  // 6. Determinar tipo de dispositivo (mobile o desktop)
  const device_type = /Mobi|Android|iPhone/i.test(ua) ? 'mobile' : 'desktop';

  // 7. Generar un identificador único para el dispositivo
  const device_id = typeof crypto !== 'undefined' && crypto.randomUUID ? crypto.randomUUID() : Math.random().toString(36).substring(2);

  // 8. Construir y retornar el objeto DeviceInfo completo
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
