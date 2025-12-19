'use server';

import { auth } from '@/lib/auth';
import type { ApplicationItem } from '@/types/applications';
import { fn_update_keycloak_client } from '@/actions/keycloak/clients/fn_update_keycloak_client';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UpdateApplicationInput {
  name?: string;
  client_secret?: string;
  domain?: string;
  logo?: string;
  description?: string;
  status?: 'active' | 'inactive' | 'development';
  keycloak_id?: string | null; // Necesario para actualizar Keycloak
}

export const fn_update_application = async (id: string, input: UpdateApplicationInput): Promise<ApplicationItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // 1. Actualizar en Keycloak si tiene keycloak_id
    if (input.keycloak_id) {
      try {
        await fn_update_keycloak_client({
          keycloakId: input.keycloak_id,
          name: input.name,
          description: input.description,
          rootUrl: input.domain,
          enabled: input.status === 'active',
        });
        console.log('Cliente actualizado en Keycloak');
      } catch (kcError) {
        console.error('Error actualizando en Keycloak (continuando con backend):', kcError);
      }
    }

    // 2. Actualizar en el backend
    const res = await fetch(`${API_BASE_URL}/api/applications/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al actualizar aplicación: ${res.statusText}`);
    }

    const data: ApplicationItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_update_application:', err);
    throw err;
  }
};