'use server';

import { ModulesListResponse } from '@/types/modules';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_modules = async (page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<ModulesListResponse> => {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/modules?${query}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener módulos: ${res.statusText}`);

    const data: ModulesListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_modules:', err);
    throw err;
  }
};
