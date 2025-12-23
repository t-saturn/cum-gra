'use server';

import { auth } from '@/lib/auth';
import type { ModuleItem } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateModuleInput {
  item?: string;
  name: string;
  route: string;
  icon?: string;
  parent_id?: string;
  application_id?: string;
  sort_order?: number;
  status?: 'active' | 'inactive';
}

export const fn_create_module = async (input: CreateModuleInput): Promise<ModuleItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/modules`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear módulo: ${res.statusText}`);
    }

    const data: ModuleItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_module:', err);
    throw err;
  }
};