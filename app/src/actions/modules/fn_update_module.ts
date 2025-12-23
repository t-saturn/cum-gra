'use server';

import { auth } from '@/lib/auth';
import type { ModuleItem } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UpdateModuleInput {
  item?: string;
  name?: string;
  route?: string;
  icon?: string;
  parent_id?: string | null;
  sort_order?: number;
  status?: 'active' | 'inactive';
}

export const fn_update_module = async (id: string, input: UpdateModuleInput): Promise<ModuleItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/modules/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al actualizar módulo: ${res.statusText}`);
    }

    const data: ModuleItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_update_module:', err);
    throw err;
  }
};