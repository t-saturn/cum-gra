'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface DeleteAssignmentResponse {
  message: string;
}

export const fn_delete_assignment = async (id: string): Promise<DeleteAssignmentResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/user-application-roles/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al eliminar asignación: ${res.statusText}`);
    }

    const data: DeleteAssignmentResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_delete_assignment:', err);
    throw err;
  }
};