'use server';
import { auth } from '@/lib/auth';
import type { UserRestrictionsStatsResponse } from '@/types/user-restrictions';
const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_user_restrictions_stats = async (): Promise<UserRestrictionsStatsResponse> => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/user-restrictions/stats`, {
    headers: { Authorization: `Bearer ${session.accessToken}` },
    cache: 'no-store',
  });

  if (!res.ok) throw new Error('Error al obtener estadísticas');
  return await res.json();
};