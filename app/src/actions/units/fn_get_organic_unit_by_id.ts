'use server';

import { auth } from '@/lib/auth';
import type { OrganicUnitItemDTO } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_organic_unit_by_id = async (id: string): Promise<OrganicUnitItemDTO> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener unidad orgánica: ${res.statusText}`);
    }

    const data: OrganicUnitItemDTO = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_organic_unit_by_id:', err);
    throw err;
  }
};