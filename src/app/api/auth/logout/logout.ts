import { NextRequest, NextResponse } from 'next/server';
import { getDeviceInfo } from '@/helpers/device-info';

export async function POST(req: NextRequest) {
  // 1. Extract session_id & logout_type from body
  const { session_id, logout_type } = await req.json();

  // 2. Read the access token cookie
  const token = req.cookies.get('access_token')?.value;
  if (!token) {
    return NextResponse.json(
      {
        success: false,
        message: 'no_token',
        data: null,
        error: { code: 'NO_TOKEN', details: 'No access token cookie found' },
      },
      { status: 401 },
    );
  }

  console.log(token);

  // 4. Build logout payload
  const payload = {
    token,
    session_id,
    logout_type,
  };

  // 5. Proxy to your Go backend
  const backendRes = await fetch(`${process.env.NEXT_PUBLIC_API_BASE}/auth/logout`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(payload),
  });

  // 6. Read backend response
  const backendJson = await backendRes.json();

  // 7. Build NextResponse with same envelope & status
  const response = NextResponse.json(backendJson, {
    status: backendRes.status,
  });

  // 8. If logout succeeded, clear the cookies
  if (backendRes.ok) {
    response.cookies.set('access_token', '', { maxAge: 0, path: '/' });
    response.cookies.set('refresh_token', '', { maxAge: 0, path: '/' });
  }

  return response;
}
