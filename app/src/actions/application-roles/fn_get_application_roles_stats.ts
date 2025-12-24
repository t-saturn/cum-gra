'use server';

import { auth } from '@/lib/auth';
import type { ApplicationRolesStatsResponse } from '@/types/application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_application_roles_stats = async (): Promise<ApplicationRolesStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/application-roles/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener estadísticas de roles: ${res.statusText}`);
    }

    const data: ApplicationRolesStatsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_application_roles_stats:', err);
    throw err;
  }
};