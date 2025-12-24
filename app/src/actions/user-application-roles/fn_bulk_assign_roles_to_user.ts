'use server';

import { auth } from '@/lib/auth';
import type { BulkAssignResponse } from '@/types/user-application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface BulkAssignRolesToUserInput {
  user_id: string;
  application_id: string;
  role_ids: string[];
}

export const fn_bulk_assign_roles_to_user = async (input: BulkAssignRolesToUserInput): Promise<BulkAssignResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles/bulk-assign-roles-to-user`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al asignar roles: ${res.statusText}`);
    }

    const data: BulkAssignResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_bulk_assign_roles_to_user:', err);
    throw err;
  }
};