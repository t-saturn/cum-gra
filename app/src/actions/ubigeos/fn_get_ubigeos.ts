'use server';

import { auth } from '@/lib/auth';
import type { UbigeosListResponse } from '@/types/ubigeos';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_ubigeos = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: {
    department?: string;
    province?: string;
    district?: string;
  }
): Promise<UbigeosListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.department) params.append('department', filters.department);
    if (filters?.province) params.append('province', filters.province);
    if (filters?.district) params.append('district', filters.district);

    const res = await fetch(`${API_BASE_URL}/api/ubigeos?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener ubigeos: ${res.statusText}`);
    }

    const data: UbigeosListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_ubigeos:', err);
    throw err;
  }
};