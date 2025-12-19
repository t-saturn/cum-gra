'use server';

import { auth } from '@/lib/auth';
import { fn_enable_keycloak_user } from '@/actions/keycloak/users/fn_enable_keycloak_user';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface RestoreUserResponse {
  message: string;
}

export const fn_restore_user = async (
  id: string,
  keycloakId?: string | null
): Promise<RestoreUserResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // 1. Habilitar en Keycloak si tiene keycloak_id
    if (keycloakId) {
      try {
        await fn_enable_keycloak_user(keycloakId);
        console.log('Usuario habilitado en Keycloak');
      } catch (kcError) {
        console.error('Error habilitando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Restaurar en el backend
    const res = await fetch(`${API_BASE_URL}/api/users/${id}/restore`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al restaurar usuario: ${res.statusText}`);
    }

    const data: RestoreUserResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_restore_user:', err);
    throw err;
  }
};