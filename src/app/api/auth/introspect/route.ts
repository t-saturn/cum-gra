import { NextRequest, NextResponse } from 'next/server';

export const runtime = 'nodejs'; // Permite usar Buffer sin problemas en Edge/Node
export const dynamic = 'force-dynamic'; // Evita caché en rutas app (introspección es dinámica)

// GET /api/auth/introspect?to=/ruta/destino
// Función: Verifica sesión llamando al backend de introspección.
// - Si hay sesión válida → redirige al destino del front (`to`, relativo) o /home por defecto.
// - Si NO hay sesión → redirige al Auth UI con `callback_url` (base64) para volver post-login.
function getOrigin(req: NextRequest) {
  // 1. Resolver origen real (http/https + host) soportando reverse proxy
  const proto = req.headers.get('x-forwarded-proto') ?? 'http';
  const host = req.headers.get('host') ?? '10.10.10.43:5556';
  return `${proto}://${host}`;
}

function toBase64(text: string) {
  // 2. Helper para codificar callback_url en base64 (URL-safe lo maneja el backend)
  return Buffer.from(text, 'utf8').toString('base64');
}

export async function GET(req: NextRequest) {
  // 3. Calcular `origin` del front y parsear la URL entrante
  const origin = getOrigin(req);
  const url = new URL(req.url);

  // 4. Resolver destino final del front:
  //    - `to` relativo (p.ej., /home). Si falta, usar /home por defecto.
  //    - Construir URL absoluta segura contra host actual.
  const to = url.searchParams.get('to') ?? '/home';
  const targetUrl = new URL(to, origin).toString();

  // 5. Preparar `callback_url` en base64 (el backend lo decodifica/valida)
  const callbackB64 = toBase64(targetUrl);

  // 6. Construir URL del backend /auth/introspect con redirect y callback_url (base64)
  const backend = new URL('http://10.10.10.43:5555/auth/introspect');
  backend.searchParams.set('redirect', 'true');
  backend.searchParams.set('callback_url', callbackB64);

  try {
    // 7. Hacer fetch al backend reenviando cookies del cliente
    //    - redirect: 'manual' para capturar 3xx
    //    - headers.Cookie: pasar cookies actuales hacia el backend
    const res = await fetch(backend.toString(), {
      method: 'GET',
      headers: {
        Cookie: req.headers.get('cookie') ?? '',
        Accept: 'application/json',
      },
      redirect: 'manual',
    });

    // 8. Si el backend manda redirección (3xx), respetarla (p.ej., hacia Auth UI)
    if (res.status >= 300 && res.status < 400) {
      const location = res.headers.get('location');
      if (location) {
        return NextResponse.redirect(location, { status: 302 });
      }
    }

    // 9. Intentar leer JSON del backend (contrato propio)
    let data = null;
    try {
      data = await res.json();
    } catch {
      // Si no hay JSON, asumimos que no se validó sesión
    }

    // 10. Si el backend confirma sesión válida → redirigir al destino del front
    if (data?.success === true) {
      return NextResponse.redirect(targetUrl, { status: 302 });
    }

    // 11. Sin sesión → redirigir al Auth UI con callback_url (base64) para volver post-login
    const authUI = `http://10.10.10.43:6160/auth/login?callback_url=${encodeURIComponent(callbackB64)}`;
    return NextResponse.redirect(authUI, { status: 302 });
  } catch {
    // 12. Error de red/infra → fallback al Auth UI con callback
    const authUI = `http://10.10.10.43:6160/auth/login?callback_url=${encodeURIComponent(callbackB64)}`;
    return NextResponse.redirect(authUI, { status: 302 });
  }
}
