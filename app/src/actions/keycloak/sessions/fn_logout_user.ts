'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export async function fn_logout_user(userId: string): Promise<void> {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        // Cerrar todas las sesiones del usuario
        await client.users.logout({ id: userId });
      },
      'Failed to logout user'
    );
  } catch (error) {
    console.error('Error logging out user:', error);
    throw error;
  }
}