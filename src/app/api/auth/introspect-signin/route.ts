import { NextRequest, NextResponse } from 'next/server';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://10.10.10.43:5555';

// GET /api/auth/introspect-signin
// Función: Reenvía cookies de sesión al gateway (/auth/introspect-signin) y controla la respuesta.
// - Si el backend responde con redirect (3xx) → propaga el Location al navegador.
// - Si el backend responde con JSON → devuelve el JSON tal cual al cliente.
// - Siempre propaga los Set-Cookie que lleguen del backend.
export const GET = async (req: NextRequest) => {
  const url = new URL(req.url);

  // 1. Leer parámetros redirect y callback_url (base64 opcional)
  const redirect = url.searchParams.get('redirect') ?? 'true';
  const callbackB64 = url.searchParams.get('callback_url') ?? '';

  // 2. Construir la URL del gateway con query params (redirect y callback_url)
  const gwURL = new URL(`${API_BASE}/auth/introspect-signin`);
  gwURL.searchParams.set('redirect', redirect);
  if (callbackB64) gwURL.searchParams.set('callback_url', callbackB64);

  // 3. Reenviar SOLO las cookies relevantes (access_token, refresh_token, session_id)
  const pick = (name: string) => req.cookies.get(name)?.value;
  const fwdCookies = [
    ['access_token', pick('access_token')],
    ['refresh_token', pick('refresh_token')],
    ['session_id', pick('session_id')],
  ]
    .filter(([, v]) => !!v) // quitar nulos/vacíos
    .map(([k, v]) => `${k}=${v}`) // formatear como cookie
    .join('; ');

  // 4. Llamada al backend (introspect-signin)
  //    - redirect: 'manual' → capturar 3xx
  //    - credentials: 'include' → mantener cookies
  //    - cache: 'no-store' → evitar caché en auth
  const backendRes = await fetch(gwURL.toString(), {
    method: 'GET',
    headers: { 'Content-Type': 'application/json', ...(fwdCookies ? { cookie: fwdCookies } : {}) },
    redirect: 'manual',
    credentials: 'include',
    cache: 'no-store',
  });

  // 5. Propagar Set-Cookie del backend al cliente
  const headers = new Headers();
  const setCookieHeader = backendRes.headers.get('set-cookie');
  if (setCookieHeader) {
    // dividir cookies múltiples en un solo header
    const cookies = setCookieHeader.split(/,(?=[^ ]*=)/g);
    for (const c of cookies) headers.append('set-cookie', c.trim());
  }

  // 6. Manejo de redirecciones: si backend devuelve 3xx → reenviar Location
  if (backendRes.status >= 300 && backendRes.status < 400) {
    const location = backendRes.headers.get('location') || '/';
    headers.set('location', location);
    return new NextResponse(null, { status: backendRes.status, headers });
  }

  // 7. Caso normal (sin redirect): devolver JSON de backend tal cual
  let json: any;
  try {
    json = await backendRes.json();
  } catch {
    json = { success: false };
  }

  // 8. Responder al cliente con el mismo status y cookies propagadas
  return new NextResponse(JSON.stringify(json), { status: backendRes.status, headers });
};
