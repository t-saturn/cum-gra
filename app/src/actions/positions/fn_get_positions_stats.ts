'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface PositionsStatsResponse {
  total_positions: number;
  active_positions: number;
  deleted_positions: number;
  assigned_employees: number;
}

export const fn_get_positions_stats = async (): Promise<PositionsStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/positions/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas de posiciones: ${res.statusText}`);
    }

    const data: PositionsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_positions_stats:', err);
    throw err;
  }
};