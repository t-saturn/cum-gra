'use server';

import { RolesAppsStatsResponse } from '@/types/roles_app';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_roles_apps_stats = async (): Promise<RolesAppsStatsResponse> => {
  try {
    const res = await fetch(`${API_BASE_URL}/roles-app/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener roles-app stats: ${res.status} ${res.statusText}`);
    }

    const data: RolesAppsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_roles_apps_stats:', err);
    throw err;
  }
};
