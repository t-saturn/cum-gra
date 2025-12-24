'use server';

import { auth } from '@/lib/auth';
import type { UserRestrictionsListResponse } from '@/types/user-restrictions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_user_restrictions = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: {
    user_id?: string;
    application_id?: string;
    is_deleted?: boolean;
  }
): Promise<UserRestrictionsListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) throw new Error('No hay sesión activa');

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.user_id) params.append('user_id', filters.user_id);
    if (filters?.application_id) params.append('application_id', filters.application_id);
    if (filters?.is_deleted) params.append('is_deleted', 'true');

    const res = await fetch(`${API_BASE_URL}/api/user-restrictions?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener restricciones: ${res.statusText}`);
    }

    const data: UserRestrictionsListResponse = await res.json();
    console.log('Restricciones obtenidas:', data);
    return data;
  } catch (err) {
    console.error('Error en fn_get_modules:', err);
    throw err;
  }
};