'use server';

import { auth } from '@/lib/auth';
import type { OrganicUnitsListResponse } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_organic_units = async (
  page: number = 1,
  pageSize: number = 10,
  isDeleted: boolean = false,
  parentId?: string | null
): Promise<OrganicUnitsListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (isDeleted) params.append('is_deleted', 'true');
    if (parentId === null) {
      params.append('parent_id', 'null');
    } else if (parentId) {
      params.append('parent_id', parentId);
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener unidades orgánicas: ${res.statusText}`);
    }

    const data: OrganicUnitsListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_organic_units:', err);
    throw err;
  }
};