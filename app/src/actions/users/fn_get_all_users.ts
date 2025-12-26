'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UserSelectItem {
  id: string;
  email: string;
  dni: string;
  first_name?: string;
  last_name?: string;
  status: string;
}

export const fn_get_all_users = async (
  onlyActive: boolean = true
): Promise<UserSelectItem[]> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams();
    if (!onlyActive) params.append('only_active', 'false');

    const url = `${API_BASE_URL}/api/users/all${params.toString() ? '?' + params.toString() : ''}`;

    const res = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener usuarios: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_all_users:', err);
    throw err;
  }
};