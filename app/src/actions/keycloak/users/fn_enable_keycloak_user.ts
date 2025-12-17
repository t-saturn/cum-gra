'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_enable_keycloak_user(keycloakId: string) {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        await client.users.update(
          { id: keycloakId },
          { enabled: true }
        );
      },
      'Failed to enable Keycloak user'
    );

    return { success: true };
  } catch (error) {
    console.error('Error enabling Keycloak user:', error);
    throw error;
  }
}