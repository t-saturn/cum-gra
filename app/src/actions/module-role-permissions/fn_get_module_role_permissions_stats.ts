'use server';

import { auth } from '@/lib/auth';
import type { ModuleRolePermissionsStatsResponse } from '@/types/module-role-permissions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_module_role_permissions_stats = async (): Promise<ModuleRolePermissionsStatsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/module-role-permissions/stats`, {
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

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_module_role_permissions_stats:', err);
    throw err;
  }
};