'use server';

import { RolesAppsResponse } from '@/types/roles_app';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_roles_apps = async (page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<RolesAppsResponse> => {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/roles-app?${query}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener roles-app: ${res.status} ${res.statusText}`);
    }

    const data: RolesAppsResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_roles_apps:', err);
    throw err;
  }
};
