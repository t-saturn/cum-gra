'use server';

import { auth } from '@/lib/auth';
import type { UbigeosStatsResponse } from '@/types/ubigeos';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_ubigeos_stats = async (): Promise<UbigeosStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/ubigeos/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas de ubigeos: ${res.statusText}`);
    }

    const data: UbigeosStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_ubigeos_stats:', err);
    throw err;
  }
};