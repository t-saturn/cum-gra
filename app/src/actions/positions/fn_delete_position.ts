'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface DeletePositionResponse {
  message: string;
}

export const fn_delete_position = async (id: string): Promise<DeletePositionResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/positions/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al eliminar posición: ${res.statusText}`);
    }

    const data: DeletePositionResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_delete_position:', err);
    throw err;
  }
};