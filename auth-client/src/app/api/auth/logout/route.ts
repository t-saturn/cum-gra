import { NextRequest, NextResponse } from 'next/server';

export async function GET(req: NextRequest) {
  // 1. Read logout_type from query
  const logout_type = req.nextUrl.searchParams.get('logout_type');
  if (!logout_type) {
    return NextResponse.json({ success: false, message: 'missing logout_type' }, { status: 400 });
  }

  // 2. Read incoming browser cookies
  const cookieHeader = req.headers.get('cookie') || '';
  if (!cookieHeader.includes('access_token')) {
    return NextResponse.json({ success: false, message: 'no_token', error: { code: 'NO_TOKEN', details: 'No access token cookie' } }, { status: 401 });
  }

  // 3. Proxy GET with cookies to your Go backend
  const backendRes = await fetch(`${process.env.NEXT_PUBLIC_API_BASE}/auth/logout?logout_type=${encodeURIComponent(logout_type)}`, {
    method: 'GET',
    headers: {
      Cookie: cookieHeader,
    },
  });

  // 4. Read backend JSON
  const backendJson = await backendRes.json();

  // 5. Build NextResponse
  const response = NextResponse.json(backendJson, { status: backendRes.status });

  // 6. Forward any Set-Cookie header(s) from backend
  const sc = backendRes.headers.get('set-cookie');
  if (sc) {
    response.headers.set('set-cookie', sc);
  }

  // Clear cookies by expiring them
  response.cookies.set('access_token', '', {
    httpOnly: true,
    secure: process.env.NODE_ENV === 'production',
    sameSite: 'strict', // lowercase!
    path: '/',
    maxAge: 0,
  });
  response.cookies.set('refresh_token', '', {
    httpOnly: true,
    secure: process.env.NODE_ENV === 'production',
    sameSite: 'strict', // lowercase!
    path: '/',
    maxAge: 0,
  });

  return response;
}
