'use server';

import { auth } from '@/lib/auth';
import type { ApplicationsStatsResponse } from '@/types/applications';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8000';

export const fn_get_applications_stats = async (): Promise<ApplicationsStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/applications/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas: ${res.statusText}`);
    }

    const data: ApplicationsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_applications_stats:', err);
    throw err;
  }
};