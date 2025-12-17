'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_disable_keycloak_user(keycloakId: string) {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        await client.users.update(
          { id: keycloakId },
          { enabled: false }
        );
      },
      'Failed to disable Keycloak user'
    );

    return { success: true };
  } catch (error) {
    console.error('Error disabling Keycloak user:', error);
    throw error;
  }
}