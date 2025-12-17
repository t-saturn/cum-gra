'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export interface UpdateKeycloakUserInput {
  keycloakId: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  enabled?: boolean;
}

export async function fn_update_keycloak_user(input: UpdateKeycloakUserInput) {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        const updateData: any = {};
        
        if (input.email !== undefined) updateData.email = input.email;
        if (input.firstName !== undefined) updateData.firstName = input.firstName;
        if (input.lastName !== undefined) updateData.lastName = input.lastName;
        if (input.enabled !== undefined) updateData.enabled = input.enabled;

        await client.users.update({ id: input.keycloakId }, updateData);
      },
      'Failed to update Keycloak user'
    );

    return { success: true };
  } catch (error) {
    console.error('Error updating Keycloak user:', error);
    throw error;
  }
}