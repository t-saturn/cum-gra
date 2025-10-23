'use server';

import type { RolesAssignmentsStatsResponse } from '@/types/roles_assignments';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_roles_assignments_stats(): Promise<RolesAssignmentsStatsResponse> {
  try {
    const res = await fetch(`${API_BASE_URL}/roles-assignments/stats`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener roles-assignments stats: ${res.status} ${res.statusText}`);
    }

    const data: RolesAssignmentsStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_roles_assignments_stats:', err);
    throw err;
  }
}
