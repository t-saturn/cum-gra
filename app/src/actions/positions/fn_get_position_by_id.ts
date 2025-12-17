'use server';

import { auth } from '@/lib/auth';
import type { StructuralPositionItem } from '@/types/structural_positions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_position_by_id = async (id: string): Promise<StructuralPositionItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/positions/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener posición: ${res.statusText}`);
    }

    const data: StructuralPositionItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_position_by_id:', err);
    throw err;
  }
};