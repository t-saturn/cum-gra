'use server';

import { auth } from '@/lib/auth';
import type { ModulesListResponse } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_modules = async (
  page: number = 1,
  pageSize: number = 10,
  filters?: {
    application_id?: string;
    is_deleted?: boolean;
  }
): Promise<ModulesListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (filters?.application_id) params.append('application_id', filters.application_id);
    if (filters?.is_deleted) params.append('is_deleted', 'true');

    const res = await fetch(`${API_BASE_URL}/api/modules?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener módulos: ${res.statusText}`);
    }

    const data: ModulesListResponse = await res.json();
    console.log('Módulos obtenidos:', data);
    
    return data;
  } catch (err) {
    console.error('Error en fn_get_modules:', err);
    throw err;
  }
};