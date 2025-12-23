'use server';

import { auth } from '@/lib/auth';
import type { ModuleItem } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_module_by_id = async (id: string): Promise<ModuleItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/modules/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener módulo: ${res.statusText}`);
    }

    const data: ModuleItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_module_by_id:', err);
    throw err;
  }
};