'use server';

import { ApplicationsStatsResponse } from '@/types/applications';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_applications_stats = async (): Promise<ApplicationsStatsResponse> => {
  try {
    const res = await fetch(`${API_BASE_URL}/applications/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener estadísticas de aplicaciones: ${res.statusText}`);

    const data: ApplicationsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_applications_stats:', err);
    throw err;
  }
};
