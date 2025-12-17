'use server';

import { auth } from '@/lib/auth';
import { fn_disable_keycloak_client } from '@/actions/keycloak/clients/fn_disable_keycloak_client';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface DeleteApplicationResponse {
  message: string;
}

export const fn_delete_application = async (
  id: string, 
  keycloakId?: string | null
): Promise<DeleteApplicationResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // 1. Deshabilitar en Keycloak si tiene keycloak_id
    if (keycloakId) {
      try {
        await fn_disable_keycloak_client(keycloakId);
        console.log('Cliente deshabilitado en Keycloak');
      } catch (kcError) {
        console.error('Error deshabilitando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Soft delete en el backend
    const res = await fetch(`${API_BASE_URL}/api/applications/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al eliminar aplicación: ${res.statusText}`);
    }

    const data: DeleteApplicationResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_delete_application:', err);
    throw err;
  }
};