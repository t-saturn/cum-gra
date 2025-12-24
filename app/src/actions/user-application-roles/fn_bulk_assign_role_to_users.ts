'use server';

import { auth } from '@/lib/auth';
import type { BulkAssignResponse } from '@/types/user-application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface BulkAssignRoleToUsersInput {
  user_ids: string[];
  application_id: string;
  application_role_id: string;
}

export const fn_bulk_assign_role_to_users = async (input: BulkAssignRoleToUsersInput): Promise<BulkAssignResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles/bulk-assign-role-to-users`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al asignar rol a usuarios: ${res.statusText}`);
    }

    const data: BulkAssignResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_bulk_assign_role_to_users:', err);
    throw err;
  }
};