import { NextRequest, NextResponse } from 'next/server';
import { getDeviceInfo } from '@/helpers/device-info';

export async function POST(req: NextRequest) {
  // 1. Parse incoming JSON
  const { email, password } = await req.json();

  // 2. Gather device_info (runs in Node.js on the server)
  const device_info = await getDeviceInfo();

  // 3. Build full payload
  const payload = {
    email,
    password,
    application_id: 'app-test-1',
    device_info,
  };

  // 4. Proxy to your Go backend
  const backendRes = await fetch(`${process.env.NEXT_PUBLIC_API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    // Include cookies from client if needed:
    credentials: 'include',
    body: JSON.stringify(payload),
  });

  // 5. Read backend response body
  const backendJson = await backendRes.json();

  console.log(backendJson);
  console.log(backendRes.status);

  // 6. Create NextResponse with the same body & status
  const response = NextResponse.json(backendJson, {
    status: backendRes.status,
  });

  // 7. Forward any Set-Cookie headers
  const cookies = backendRes.headers.get('set-cookie');
  if (cookies) {
    response.headers.set('set-cookie', cookies);
  }

  return response;
}
