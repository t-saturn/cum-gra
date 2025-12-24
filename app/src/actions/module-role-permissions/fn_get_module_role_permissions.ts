'use server';

import { auth } from '@/lib/auth';
import type { ModuleRolePermissionsListResponse, ModuleRolePermissionFilters } from '@/types/module-role-permissions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_module_role_permissions = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: ModuleRolePermissionFilters
): Promise<ModuleRolePermissionsListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.module_id) params.append('module_id', filters.module_id);
    if (filters?.role_id) params.append('role_id', filters.role_id);
    if (filters?.is_deleted) params.append('is_deleted', 'true');

    const res = await fetch(`${API_BASE_URL}/api/module-role-permissions?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener permisos: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_module_role_permissions:', err);
    throw err;
  }
};