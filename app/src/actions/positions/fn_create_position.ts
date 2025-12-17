'use server';

import { auth } from '@/lib/auth';
import type { StructuralPositionItem } from '@/types/structural_positions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreatePositionInput {
  name: string;
  code: string;
  level?: number;
  description?: string;
  is_active?: boolean;
  cod_car_sgd?: string;
}

export const fn_create_position = async (input: CreatePositionInput): Promise<StructuralPositionItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/positions`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear posición: ${res.statusText}`);
    }

    const data: StructuralPositionItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_position:', err);
    throw err;
  }
};