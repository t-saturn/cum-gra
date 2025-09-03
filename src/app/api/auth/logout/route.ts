import { NextRequest, NextResponse } from 'next/server';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://10.10.10.43:5555';

// GET /api/auth/logout
// Función: Delegar el logout al gateway (/auth/logout), reenviando cookies y controlando la redirección o la respuesta JSON.
// - Si el backend responde con 3xx → propaga la redirección y limpia cookies en el front.
// - Si responde JSON → lo devuelve tal cual y limpia cookies en éxito.
export async function GET(req: NextRequest) {
  const url = new URL(req.url);

  // 1. Leer parámetros redirect y logout_type (con default)
  const redirect = url.searchParams.get('redirect') ?? 'true';
  const logoutType = url.searchParams.get('logout_type') ?? 'user_logout';

  // 2. Resolver callback_url (base64) → si no viene, usar la raíz del front actual
  const defaultTarget = `${url.origin}/`;
  const callbackB64 = url.searchParams.get('callback_url') ?? Buffer.from(defaultTarget, 'utf-8').toString('base64');

  // 3. Construir URL al gateway con los parámetros de logout
  const gwURL = new URL(`${API_BASE}/auth/logout`);
  gwURL.searchParams.set('redirect', redirect);
  gwURL.searchParams.set('logout_type', logoutType);
  gwURL.searchParams.set('callback_url', callbackB64);

  // 4. Reenviar SOLO las cookies relevantes (access_token, refresh_token, session_id)
  const pick = (name: string) => req.cookies.get(name)?.value;
  const fwdCookies = [
    ['access_token', pick('access_token')],
    ['refresh_token', pick('refresh_token')],
    ['session_id', pick('session_id')],
  ]
    .filter(([, v]) => !!v)
    .map(([k, v]) => `${k}=${v}`)
    .join('; ');

  // 5. Llamada al backend /auth/logout
  const backendRes = await fetch(gwURL.toString(), {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...(fwdCookies ? { cookie: fwdCookies } : {}),
    },
    redirect: 'manual', // capturar 3xx
    credentials: 'include',
    cache: 'no-store',
  });

  // 6. Helper para limpiar cookies locales en la respuesta al cliente
  const clearAuthCookies = (res: NextResponse) => {
    const opts = { path: '/', expires: new Date(0) };
    res.cookies.set('access_token', '', opts);
    res.cookies.set('refresh_token', '', opts);
    res.cookies.set('session_id', '', opts);
  };

  // 7. Helper para propagar Set-Cookie del backend (si hace algo extra)
  const propagateSetCookies = (res: NextResponse) => {
    const setCookieHeader = backendRes.headers.get('set-cookie');
    if (setCookieHeader) {
      const cookies = setCookieHeader.split(/,(?=[^ ]*=)/g);
      for (const c of cookies) res.headers.append('set-cookie', c.trim());
    }
  };

  // 8. Manejo de redirecciones:
  //    → Si backend responde con 3xx → propagar Location, limpiar cookies y retornar
  if (backendRes.status >= 300 && backendRes.status < 400) {
    const location = backendRes.headers.get('location') || '/';
    const out = NextResponse.redirect(location, backendRes.status);
    clearAuthCookies(out);
    propagateSetCookies(out);
    return out;
  }

  // 9. Caso normal (sin redirect): intentar parsear JSON
  let json;
  try {
    json = await backendRes.json();
  } catch {
    json = { success: false };
  }

  // 10. Responder al cliente con el JSON y status del backend
  const out = NextResponse.json(json, { status: backendRes.status });

  // 11. Si logout fue exitoso → limpiar cookies locales
  if (json?.success === true) {
    clearAuthCookies(out);
  }

  // 12. Propagar cookies extra del backend
  propagateSetCookies(out);
  return out;
}
