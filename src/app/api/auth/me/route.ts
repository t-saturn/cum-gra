import { NextRequest, NextResponse } from 'next/server';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://10.10.10.43:5555';

// GET /api/auth/me
// Función: Obtener información de la sesión/usuario desde el API Gateway (/auth/me).
// - Reenvía cookies (access_token, refresh_token, session_id).
// - Propaga cookies del backend si son renovadas.
// - Devuelve siempre JSON con el mismo status que el backend.
export const GET = async (req: NextRequest) => {
  const url = new URL(req.url);

  // 1. Leer parámetro client_id o usar un valor por defecto
  const clientId = url.searchParams.get('client_id') || 'cum';

  // 2. Construir URL al API Gateway
  const gwURL = new URL(`${API_BASE}/auth/me`);
  gwURL.searchParams.set('client_id', clientId);

  // 3. Reenviar SOLO las cookies necesarias
  const pick = (name: string) => req.cookies.get(name)?.value;
  const fwdCookies = [
    ['access_token', pick('access_token')],
    ['refresh_token', pick('refresh_token')],
    ['session_id', pick('session_id')],
  ]
    .filter(([, v]) => !!v) // quitar las que estén vacías
    .map(([k, v]) => `${k}=${v}`) // formatear cookie
    .join('; '); // unir en header

  // 4. Llamar al backend con las cookies reenviadas
  const backendRes = await fetch(gwURL.toString(), {
    method: 'GET',
    headers: { 'Content-Type': 'application/json', ...(fwdCookies ? { cookie: fwdCookies } : {}) },
    redirect: 'manual', // no seguir redirecciones
    credentials: 'include', // incluir credenciales
    cache: 'no-store', // evitar cacheo
  });

  // 5. Propagar Set-Cookie del backend (ej: refresh de tokens)
  const outHeaders = new Headers();
  const setCookieHeader = backendRes.headers.get('set-cookie');
  if (setCookieHeader) {
    // dividir múltiples cookies en el mismo header
    const cookies = setCookieHeader.split(/,(?=[^ ]*=)/g);
    for (const c of cookies) outHeaders.append('set-cookie', c.trim());
  }

  // 6. Intentar parsear JSON de respuesta, fallback en caso de error
  let json;
  // let invalidJson = false;
  try {
    json = await backendRes.json();
  } catch {
    json = { success: false, message: 'Invalid JSON from gateway' };
    // invalidJson = true; --
  }

  // 7. Contruir la respuesta para el cliente
  const res = new NextResponse(JSON.stringify(json), {
    status: backendRes.status,
    headers: outHeaders,
  });

  // 8. limpiar cookies si falla la petición
  // Consideramos falla si: status >= 400, o success === false, o JSON inválido
  // if (backendRes.status >= 400 || json?.success === false || invalidJson) { --
  //   const opts = { path: '/', expires: new Date(0) }; --
  //   res.cookies.set('access_token', '', opts); --
  //   res.cookies.set('refresh_token', '', opts); --
  //   res.cookies.set('session_id', '', opts); --
  // } --

  // 9. Devolver al cliente JSON + status del backend + cookies propagadas
  return res;
};
