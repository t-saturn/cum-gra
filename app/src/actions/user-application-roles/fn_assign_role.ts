'use server';

import { auth } from '@/lib/auth';
import type { UserApplicationRoleItem } from '@/types/user-application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface AssignRoleInput {
  user_id: string;
  application_id: string;
  application_role_id: string;
}

export const fn_assign_role = async (input: AssignRoleInput): Promise<UserApplicationRoleItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al asignar rol: ${res.statusText}`);
    }

    const data: UserApplicationRoleItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_assign_role:', err);
    throw err;
  }
};