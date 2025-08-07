import { NextRequest, NextResponse } from 'next/server';

export async function POST(req: NextRequest) {
  const { email, password } = await req.json();

  // Llama al backend real (tu auth-service en Go, por ejemplo)
  const response = await fetch(`${process.env.API_URL}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    return NextResponse.json({ error: 'Credenciales inválidas' }, { status: 401 });
  }

  const { access_token, refresh_token } = await response.json();

  const res = NextResponse.json({ success: true });

  res.cookies.set('access_token', access_token, {
    httpOnly: true,
    secure: true,
    sameSite: 'lax',
    path: '/',
    maxAge: 60 * 60, // 1 hora
  });

  res.cookies.set('refresh_token', refresh_token, {
    httpOnly: true,
    secure: true,
    sameSite: 'lax',
    path: '/',
    maxAge: 60 * 60 * 24 * 7, // 7 días
  });

  return res;
}
