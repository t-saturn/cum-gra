'use server';

import type { RolesRestrictResponse } from '@/types/users_restrictions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_users_restrictions(page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<RolesRestrictResponse> {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/users-restrictions?${query}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener users-restrictions: ${res.status} ${res.statusText}`);
    }

    const data: RolesRestrictResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_users_restrictions:', err);
    throw err;
  }
}
