'use server';

import { auth } from '@/lib/auth';
import type { ApplicationItem } from '@/types/applications';
import { fn_create_keycloak_client } from '@/actions/keycloak/clients/fn_create_keycloak_client';
import { fn_get_keycloak_client_by_clientid } from '@/actions/keycloak/clients/fn_get_keycloak_client_by_clientid';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateApplicationInput {
  name: string;
  client_id: string;
  client_secret?: string;
  domain: string;
  logo?: string;
  description?: string;
  status?: 'active' | 'inactive' | 'development';
  sync_to_keycloak?: boolean; // Nuevo flag
}

export const fn_create_application = async (input: CreateApplicationInput): Promise<ApplicationItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    let keycloakId: string | undefined;

    // 1. Crear en Keycloak si sync_to_keycloak es true
    if (input.sync_to_keycloak !== false) { // Por defecto true
      try {
        // Verificar si ya existe en Keycloak
        const existingClient = await fn_get_keycloak_client_by_clientid(input.client_id);
        
        if (existingClient) {
          keycloakId = existingClient.id;
          console.log('Cliente ya existe en Keycloak:', keycloakId);
        } else {
          const keycloakResult = await fn_create_keycloak_client({
            clientId: input.client_id,
            name: input.name,
            description: input.description,
            rootUrl: input.domain,
            enabled: input.status === 'active',
          });
          keycloakId = keycloakResult.clientId;
          console.log('Cliente creado en Keycloak:', keycloakId);
        }
      } catch (kcError) {
        console.error('Error creando en Keycloak (continuando con backend):', kcError);
        // No lanzar error, continuar con backend
      }
    }

    // 2. Crear en el backend
    const res = await fetch(`${API_BASE_URL}/api/applications`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify({
        ...input,
        keycloak_id: keycloakId, // Guardar el ID de Keycloak
      }),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear aplicación: ${res.statusText}`);
    }

    const data: ApplicationItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_application:', err);
    throw err;
  }
};