'use server';

import { getDeviceInfo } from '@/helpers/device-info';

export const loginAction = async (formData: FormData) => {
  const email = formData.get('email') as string;
  const password = formData.get('password') as string;
  const remember = formData.get('remember') === 'on';

  // recoge toda la info del dispositivo
  const device_info = await getDeviceInfo();

  const payload = {
    email,
    password,
    application_id: 'app_001',
    remember_me: remember,
    device_info,
  };

  const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(`Login failed: ${text}`);
  }

  return true;
};
