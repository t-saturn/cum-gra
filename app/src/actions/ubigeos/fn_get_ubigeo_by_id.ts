'use server';

import { auth } from '@/lib/auth';
import type { UbigeoItem } from '@/types/ubigeos';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_ubigeo_by_id = async (id: string): Promise<UbigeoItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/ubigeos/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener ubigeo: ${res.statusText}`);
    }

    const data: UbigeoItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_ubigeo_by_id:', err);
    throw err;
  }
};