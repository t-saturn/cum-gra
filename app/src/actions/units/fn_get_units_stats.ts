'use server';

import { auth } from '@/lib/auth';
import type { OrganicUnitsStatsResponse } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_units_stats = async (): Promise<OrganicUnitsStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas de unidades: ${res.statusText}`);
    }

    const data: OrganicUnitsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_units_stats:', err);
    throw err;
  }
};