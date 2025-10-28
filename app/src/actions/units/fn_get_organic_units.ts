'use server';

import { OrganicUnitsListResponse } from '@/types/units';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_organic_units = async (page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<OrganicUnitsListResponse> => {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/units?${query.toString()}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener unidades orgánicas: ${res.statusText}`);

    const data: OrganicUnitsListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_organic_units:', err);
    throw err;
  }
};
