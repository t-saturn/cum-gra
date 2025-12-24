'use server';

import { auth } from '@/lib/auth';
import type { BulkAssignPermissionsInput, BulkAssignPermissionsResponse } from '@/types/module-role-permissions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_bulk_assign_permissions = async (
  input: BulkAssignPermissionsInput
): Promise<BulkAssignPermissionsResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/module-role-permissions/bulk-assign`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}));
      throw new Error(errorData.error || `Error en asignación masiva: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_bulk_assign_permissions:', err);
    throw err;
  }
};