'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface ModuleSelectItem {
  id: string;
  name: string;
  item?: string;
  icon?: string;
  parent_id?: string;
  application_id?: string;
  status: string;
}

export const fn_get_all_modules = async (
  onlyActive: boolean = true,
  applicationId?: string
): Promise<ModuleSelectItem[]> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams();
    if (!onlyActive) params.append('only_active', 'false');
    if (applicationId) params.append('application_id', applicationId);

    const url = `${API_BASE_URL}/api/modules/all${params.toString() ? '?' + params.toString() : ''}`;

    const res = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener módulos: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_all_modules:', err);
    throw err;
  }
};