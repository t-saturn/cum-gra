'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface RevokeRoleInput {
  reason?: string;
}

export interface RevokeRoleResponse {
  message: string;
}

export const fn_revoke_role = async (id: string, input?: RevokeRoleInput): Promise<RevokeRoleResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles/${id}/revoke`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input || {}),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al revocar rol: ${res.statusText}`);
    }

    const data: RevokeRoleResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_revoke_role:', err);
    throw err;
  }
};