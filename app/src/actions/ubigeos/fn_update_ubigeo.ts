'use server';

import { auth } from '@/lib/auth';
import type { UbigeoItem } from '@/types/ubigeos';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UpdateUbigeoInput {
  ubigeo_code?: string;
  inei_code?: string;
  department?: string;
  province?: string;
  district?: string;
}

export const fn_update_ubigeo = async (id: string, input: UpdateUbigeoInput): Promise<UbigeoItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/ubigeos/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al actualizar ubigeo: ${res.statusText}`);
    }

    const data: UbigeoItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_update_ubigeo:', err);
    throw err;
  }
};