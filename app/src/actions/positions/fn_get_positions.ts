'use server';

import { auth } from '@/lib/auth';
import type { StructuralPositionsListResponse } from '@/types/structural_positions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_positions = async (
  page: number = 1,
  pageSize: number = 10,
  isDeleted: boolean = false,
  level?: number
): Promise<StructuralPositionsListResponse> => {
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
    if (level !== undefined) params.append('level', level.toString());

    const res = await fetch(`${API_BASE_URL}/api/positions?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener posiciones: ${res.statusText}`);
    }

    const data: StructuralPositionsListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_positions:', err);
    throw err;
  }
};