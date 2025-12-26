'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface ApplicationRoleSelectItem {
  id: string;
  name: string;
  description?: string;
  application_id: string;
}

export const fn_get_all_application_roles = async (
  applicationId?: string
): Promise<ApplicationRoleSelectItem[]> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams();
    if (applicationId) params.append('application_id', applicationId);

    const url = `${API_BASE_URL}/api/application-roles/all${params.toString() ? '?' + params.toString() : ''}`;

    const res = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener roles: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_get_all_application_roles:', err);
    throw err;
  }
};