'use server';

import { auth } from '@/lib/auth';
import type { UsersStatsResponse } from '@/types/users';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_user_stats = async (): Promise<UsersStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/users/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas de usuarios: ${res.statusText}`);
    }

    const data: UsersStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_user_stats:', err);
    throw err;
  }
};