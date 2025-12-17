'use server';

import { auth } from '@/lib/auth';
import type { OrganicUnitItemDTO } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UpdateOrganicUnitInput {
  name?: string;
  acronym?: string;
  brand?: string;
  description?: string;
  parent_id?: string | null;
  is_active?: boolean;
  cod_dep_sgd?: string;
}

export const fn_update_organic_unit = async (id: string, input: UpdateOrganicUnitInput): Promise<OrganicUnitItemDTO> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al actualizar unidad orgánica: ${res.statusText}`);
    }

    const data: OrganicUnitItemDTO = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_update_organic_unit:', err);
    throw err;
  }
};