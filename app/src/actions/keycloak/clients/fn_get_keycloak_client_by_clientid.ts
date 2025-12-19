'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_get_keycloak_client_by_clientid(clientId: string) {
  try {
    const result = await keycloakAdmin.execute(
      async (client) => {
        const clients = await client.clients.find({ clientId });
        return clients.length > 0 ? clients[0] : null;
      },
      'Failed to get Keycloak client'
    );

    return result;
  } catch (error) {
    console.error('Error getting Keycloak client:', error);
    return null;
  }
}