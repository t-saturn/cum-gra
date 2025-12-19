'use server';

import { auth } from '@/lib/auth';
import type { OrganicUnitItemDTO } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateOrganicUnitInput {
  name: string;
  acronym: string;
  brand?: string;
  description?: string;
  parent_id?: string;
  is_active?: boolean;
  cod_dep_sgd?: string;
}

export const fn_create_organic_unit = async (input: CreateOrganicUnitInput): Promise<OrganicUnitItemDTO> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear unidad orgánica: ${res.statusText}`);
    }

    const data: OrganicUnitItemDTO = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_organic_unit:', err);
    throw err;
  }
};