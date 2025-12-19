'use server';

import { auth } from '@/lib/auth';
import { fn_disable_keycloak_user } from '@/actions/keycloak/users/fn_disable_keycloak_user';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface DeleteUserResponse {
  message: string;
}

export const fn_delete_user = async (
  id: string,
  keycloakId?: string | null
): Promise<DeleteUserResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // 1. Deshabilitar en Keycloak si tiene keycloak_id
    if (keycloakId) {
      try {
        await fn_disable_keycloak_user(keycloakId);
        console.log('Usuario deshabilitado en Keycloak');
      } catch (kcError) {
        console.error('Error deshabilitando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Soft delete en el backend
    const res = await fetch(`${API_BASE_URL}/api/users/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al eliminar usuario: ${res.statusText}`);
    }

    const data: DeleteUserResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_delete_user:', err);
    throw err;
  }
};