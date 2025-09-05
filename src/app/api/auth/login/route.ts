import { NextRequest, NextResponse } from 'next/server';

// URL base del backend (API Gateway o servicio de autenticación)
const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:5555';

// POST /api/auth/login
// Body: { username, password, callback_url? }
// Función: Delegar el login al backend, reenviando cookies y controlando redirecciones.
export async function POST(req: NextRequest) {
  // 1. Leer y parsear el cuerpo de la petición
  const body = await req.json();

  // 2. Si no se proporciona callback_url, usar una URL por defecto del frontend (opcional - /home)
  if (!body.callback_url) {
    const FRONT_BASE = process.env.NEXT_PUBLIC_FRONT_BASE || 'http://localhost:5556';
    // body.callback_url = `${FRONT_BASE}/home`;
    body.callback_url = `${FRONT_BASE}`;
  }

  // 3. Realizar la petición al backend
  // - redirect: 'manual' → no seguir redirecciones automáticamente
  // - credentials: 'include' → permitir cookies del backend
  const backendRes = await fetch(`${API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    redirect: 'manual', // capturar 303, 302, etc.
    credentials: 'include', // reenviar/recibir cookies
    body: JSON.stringify(body),
  });

  // 4. Preparar el reenvío de cookies del backend hacia el cliente
  const headers = new Headers();
  const raw = (backendRes as any).headers?.raw?.()['set-cookie'] as string[] | undefined;

  if (raw?.length) {
    // Si el backend envía múltiples cookies
    for (const c of raw) headers.append('set-cookie', c);
  } else {
    // Si el backend envía una sola cookie
    const single = backendRes.headers.get('set-cookie');
    if (single) headers.set('set-cookie', single);
  }

  // 5. Si el backend responde con un código 3xx (redirección)
  //    → devolver { redirect } para que el frontend decida la navegación
  if (backendRes.status >= 300 && backendRes.status < 400) {
    const location = backendRes.headers.get('location') || '/';
    return NextResponse.json({ redirect: location }, { headers });
  }

  // 6. Si el backend responde con 401 (no autorizado), reenviar el error tal cual
  if (backendRes.status === 401) {
    const json = await backendRes.json().catch(() => ({
      error: { code: 'UNAUTHORIZED' },
    }));
    return new NextResponse(JSON.stringify(json), { status: 401, headers });
  }

  // 7. Fallback: procesar la respuesta estándar
  //    - Si no es redirección ni 401, devolver el JSON del backend o un mensaje genérico
  let json: any = null;
  try {
    json = await backendRes.json();
  } catch {
    json = { message: 'OK' };
  }

  // 8. Responder al cliente con el mismo status del backend y las cookies reenviadas
  return new NextResponse(JSON.stringify(json), {
    status: backendRes.status,
    headers,
  });
}
