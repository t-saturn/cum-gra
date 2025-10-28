'use server';

import type { UsersRestrictionsStatsResponse } from '@/types/users_restrictions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_users_restrictions_stats(): Promise<UsersRestrictionsStatsResponse> {
  try {
    const res = await fetch(`${API_BASE_URL}/users-restrictions/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener users-restrictions stats: ${res.status} ${res.statusText}`);

    const data: UsersRestrictionsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_users_restrictions_stats:', err);
    throw err;
  }
}
