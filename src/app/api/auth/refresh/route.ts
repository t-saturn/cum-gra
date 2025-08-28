/* eslint-disable @typescript-eslint/no-explicit-any */
import { NextRequest, NextResponse } from 'next/server';

export const runtime = 'nodejs'; // asegura acceso a headers.raw() en Node
const AUTH_BASE_URL = process.env.AUTH_BASE_URL ?? 'http://localhost:5555';

export async function GET(req: NextRequest) {
  // 1. Leer cookies requeridas
  const sessionID = req.cookies.get('session_id')?.value ?? '';
  const access = req.cookies.get('access_token')?.value ?? '';
  const refresh = req.cookies.get('refresh_token')?.value ?? '';

  if (!sessionID || !access || !refresh) {
    return NextResponse.json(
      {
        success: false,
        message: 'Faltan cookies requeridas',
        error: {
          code: 'MISSING_FIELDS',
          message: 'session_id, access_token y refresh_token son requeridos (cookies)',
          details: {
            missing_session_id: !sessionID,
            missing_access_token: !access,
            missing_refresh_token: !refresh,
          },
        },
      },
      { status: 400 },
    );
  }

  // 2. Construir header Cookie para el upstream
  const cookieHeader = [`session_id=${encodeURIComponent(sessionID)}`, `access_token=${encodeURIComponent(access)}`, `refresh_token=${encodeURIComponent(refresh)}`].join('; ');

  // 3. Hacer GET al auth-service /auth/refresh reenviando SOLO cookies
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), 10_000);

  let upstream: Response;
  try {
    upstream = await fetch(`${AUTH_BASE_URL}/auth/refresh`, {
      method: 'GET', // <-- ahora GET
      headers: {
        Accept: 'application/json',
        Cookie: cookieHeader,
        'Cache-Control': 'no-store',
      },
      redirect: 'manual',
      signal: controller.signal,
    });
  } catch (e: any) {
    clearTimeout(timer);
    return NextResponse.json(
      {
        success: false,
        message: 'No se pudo contactar al servicio de autenticación',
        error: {
          code: 'HTTP_REQUEST_ERROR',
          message: e?.message ?? String(e),
        },
      },
      { status: 502 },
    );
  } finally {
    clearTimeout(timer);
  }

  // 4. Reenviar cuerpo y status tal cual al cliente
  const text = await upstream.text();
  const contentType = upstream.headers.get('content-type') ?? 'application/json';

  const res = new NextResponse(text, {
    status: upstream.status,
    headers: {
      'content-type': contentType,
      'cache-control': 'no-store',
    },
  });

  // 5. Propagar los Set-Cookie del upstream al navegador
  const anyHeaders = upstream.headers as any;
  const raw = typeof anyHeaders.raw === 'function' ? anyHeaders.raw() : null;
  const setCookies: string[] = raw?.['set-cookie'] ?? [];

  if (setCookies.length > 0) {
    for (const sc of setCookies) {
      res.headers.append('set-cookie', sc);
    }
  } else {
    const single = upstream.headers.get('set-cookie');
    if (single) res.headers.append('set-cookie', single);
  }

  return res;
}
