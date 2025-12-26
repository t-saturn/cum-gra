'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface OrganicUnitSelectItem {
  id: string;
  name: string;
  acronym: string;
  parent_id?: string;
  is_active: boolean;
}

export const fn_get_all_organic_units = async (
  onlyActive: boolean = true
): Promise<OrganicUnitSelectItem[]> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams();
    if (!onlyActive) params.append('only_active', 'false');

    const url = `${API_BASE_URL}/api/organic-units/all${params.toString() ? '?' + params.toString() : ''}`;

    const res = await fetch(url, {
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

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_all_organic_units:', err);
    throw err;
  }
};