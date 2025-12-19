'use server';

import { auth } from '@/lib/auth';
import type { UsersListResponse } from '@/types/users';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_users = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: {
    status?: 'active' | 'suspended' | 'inactive';
    organic_unit_id?: string;
    position_id?: string;
    is_deleted?: boolean;
  }
): Promise<UsersListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.status) params.append('status', filters.status);
    if (filters?.organic_unit_id) params.append('organic_unit_id', filters.organic_unit_id);
    if (filters?.position_id) params.append('position_id', filters.position_id);
    if (filters?.is_deleted) params.append('is_deleted', 'true');

    const res = await fetch(`${API_BASE_URL}/api/users?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener usuarios: ${res.statusText}`);
    }

    const data: UsersListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_users:', err);
    throw err;
  }
};