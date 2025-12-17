'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_get_keycloak_user_by_username(username: string) {
  try {
    const result = await keycloakAdmin.execute(
      async (client) => {
        const users = await client.users.find({ username, exact: true });
        return users.length > 0 ? users[0] : null;
      },
      'Failed to get Keycloak user'
    );

    return result;
  } catch (error) {
    console.error('Error getting Keycloak user:', error);
    return null;
  }
}