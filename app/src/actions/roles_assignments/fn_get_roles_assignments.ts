'use server';

import type { RolesAssignmentsResponse } from '@/types/roles_assignments';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_roles_assignments(page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<RolesAssignmentsResponse> {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/roles-assignments?${query}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener roles-assignments: ${res.status} ${res.statusText}`);
    }

    const data: RolesAssignmentsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_roles_assignments:', err);
    throw err;
  }
}
