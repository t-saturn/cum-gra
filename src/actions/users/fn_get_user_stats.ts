'use server';

import { UsersStatsResponse } from '@/types/users';
import { revalidateTag } from 'next/cache';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function getUsersStats(): Promise<UsersStatsResponse> {
  try {
    const res = await fetch(`${API_BASE_URL}/users/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store', // o 'force-cache' si quieres caching
    });

    if (!res.ok) throw new Error(`Error al obtener estadísticas de usuarios: ${res.statusText}`);

    const data = (await res.json()) as UsersStatsResponse;
    return data;
  } catch (error) {
    console.error('Error al llamar al endpoint /users/stats:', error);
    throw error;
  }
}
