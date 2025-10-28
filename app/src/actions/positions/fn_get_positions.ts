'use server';

import { StructuralPositionsListResponse } from '@/types/structural_positions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export const fn_get_positions = async (page: number = 1, pageSize: number = 20, isDeleted: boolean = false): Promise<StructuralPositionsListResponse> => {
  try {
    const query = new URLSearchParams({
      page: String(page),
      page_size: String(pageSize),
      is_deleted: String(isDeleted),
    });

    const res = await fetch(`${API_BASE_URL}/positions?${query}`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error al obtener posiciones: ${res.statusText}`);

    const data: StructuralPositionsListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_positions:', err);
    throw err;
  }
};
