import { NextRequest, NextResponse } from 'next/server';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:5555';

// GET /api/auth/introspect
// Objetivo: Verificar sesión activa a través del API Gateway, controlando redirecciones opcionales.
export async function GET(req: NextRequest) {
  // 1. Extraer query params desde la URL recibida
  //    - redirect: controla si el API Gateway debe redirigir (true/false)
  const url = new URL(req.url);
  const redirect = url.searchParams.get('redirect') ?? 'true';

  // 2. Llamar al API Gateway → /auth/introspect
  //    - Reenviamos cookies actuales para validar la sesión.
  //    - Usamos redirect: 'manual' para capturar redirecciones sin que fetch las siga.
  const backendRes = await fetch(`${API_BASE}/auth/introspect?redirect=${redirect}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      // MUY IMPORTANTE: reenviar cookies del browser al API Gateway
      cookie: req.headers.get('cookie') ?? '',
    },
    redirect: 'manual', // para manejar 302/303 manualmente
    credentials: 'include', // por si el API Gateway setea cookies nuevas
    cache: 'no-store', // evitar respuestas cacheadas
  });

  // 3. Preparar encabezados para propagar cookies al cliente
  const headers = new Headers();
  const raw = (backendRes as any).headers?.raw?.()['set-cookie'] as string[] | undefined;

  if (raw?.length) {
    // Caso: múltiples cookies
    for (const c of raw) headers.append('set-cookie', c);
  } else {
    // Caso: solo una cookie
    const single = backendRes.headers.get('set-cookie');
    if (single) headers.set('set-cookie', single);
  }

  // 4. Si el backend responde con un código 3xx (redirección)
  //    → devolvemos un objeto { redirect } para que el cliente decida la navegación.
  if (backendRes.status >= 300 && backendRes.status < 400) {
    const location = backendRes.headers.get('location') || '/';
    return NextResponse.json({ redirect: location }, { status: 200, headers });
  }

  // 5. Si no hubo redirección → devolver el JSON de introspect tal cual.
  let json: any;
  try {
    json = await backendRes.json();
  } catch {
    json = { success: false };
  }

  // 6. Responder al cliente manteniendo el mismo status y cookies del backend.
  return new NextResponse(JSON.stringify(json), {
    status: backendRes.status,
    headers,
  });
}
