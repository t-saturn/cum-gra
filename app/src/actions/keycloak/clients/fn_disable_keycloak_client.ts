'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_disable_keycloak_client(keycloakId: string) {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        await client.clients.update(
          { id: keycloakId },
          { enabled: false }
        );
      },
      'Failed to disable Keycloak client'
    );

    return { success: true };
  } catch (error) {
    console.error('Error disabling Keycloak client:', error);
    throw error;
  }
}