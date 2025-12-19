'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export interface UpdateKeycloakClientInput {
  keycloakId: string; // ID interno de Keycloak
  name?: string;
  description?: string;
  rootUrl?: string;
  enabled?: boolean;
}

export async function fn_update_keycloak_client(input: UpdateKeycloakClientInput) {
  try {
    await keycloakAdmin.execute(
      async (client) => {
        const updateData: any = {};
        
        if (input.name !== undefined) updateData.name = input.name;
        if (input.description !== undefined) updateData.description = input.description;
        if (input.rootUrl !== undefined) {
          updateData.rootUrl = input.rootUrl;
          updateData.baseUrl = input.rootUrl;
        }
        if (input.enabled !== undefined) updateData.enabled = input.enabled;

        await client.clients.update({ id: input.keycloakId }, updateData);
      },
      'Failed to update Keycloak client'
    );

    return { success: true };
  } catch (error) {
    console.error('Error updating Keycloak client:', error);
    throw error;
  }
}