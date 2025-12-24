'use server';

import { auth } from '@/lib/auth';
import type { UserApplicationRolesListResponse } from '@/types/user-application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_user_application_roles = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: {
    user_id?: string;
    application_id?: string;
    is_revoked?: boolean;
    is_deleted?: boolean;
  }
): Promise<UserApplicationRolesListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.user_id) params.append('user_id', filters.user_id);
    if (filters?.application_id) params.append('application_id', filters.application_id);
    if (filters?.is_revoked !== undefined) params.append('is_revoked', String(filters.is_revoked));
    if (filters?.is_deleted) params.append('is_deleted', 'true');

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener asignaciones de roles: ${res.statusText}`);
    }

    const data: UserApplicationRolesListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_user_application_roles:', err);
    throw err;
  }
};