'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface DeleteUbigeoResponse {
  message: string;
}

export const fn_delete_ubigeo = async (id: string): Promise<DeleteUbigeoResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/ubigeos/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al eliminar ubigeo: ${res.statusText}`);
    }

    const data: DeleteUbigeoResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_delete_ubigeo:', err);
    throw err;
  }
};