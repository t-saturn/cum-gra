'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import { mapKeycloakClientToApplication } from '@/lib/keycloak/mappers';
import type { KeycloakApplication } from '@/types/keycloak/clients';

export interface GetClientsResult {
  success: boolean;
  data: KeycloakApplication[];
  error?: string;
}

export async function getKeycloakClients(): Promise<GetClientsResult> {
  try {
    const clients = await keycloakAdmin.execute(
      async (client) => {
        // Obtener todos los clientes del realm
        return await client.clients.find();
      },
      'Failed to fetch Keycloak clients'
    );

    // Filtrar clientes internos de Keycloak (opcional)
    const filteredClients = clients.filter(
      (c) =>
        !c.clientId?.startsWith('realm-') &&
        !c.clientId?.includes('admin-cli') &&
        !c.clientId?.includes('account') &&
        !c.clientId?.includes('broker') &&
        !c.clientId?.includes('security-admin-console')
    );

    // Mapear a tu estructura ApplicationItem
    const mappedClients = filteredClients.map(mapKeycloakClientToApplication);

    return {
      success: true,
      data: mappedClients,
    };
  } catch (error) {
    console.error('Error fetching Keycloak clients:', error);
    return {
      success: false,
      data: [],
      error: error instanceof Error ? error.message : 'Unknown error',
    };
  }
}