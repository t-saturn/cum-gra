'use server';

import { ModulesStatsResponse } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_modules_stats = async (): Promise<ModulesStatsResponse> => {
  try {
    const res = await fetch(`${API_BASE_URL}/modules/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener estadísticas de módulos: ${res.statusText}`);

    const data: ModulesStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_modules_stats:', err);
    throw err;
  }
};
