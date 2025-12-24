'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface RestoreRoleResponse {
  message: string;
}

export const fn_restore_role = async (id: string): Promise<RestoreRoleResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles/${id}/restore`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al restaurar rol: ${res.statusText}`);
    }

    const data: RestoreRoleResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_restore_role:', err);
    throw err;
  }
};