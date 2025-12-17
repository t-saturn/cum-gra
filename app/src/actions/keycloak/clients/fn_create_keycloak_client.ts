'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import type ClientRepresentation from '@keycloak/keycloak-admin-client/lib/defs/clientRepresentation';

export interface CreateKeycloakClientInput {
  clientId: string;
  name: string;
  description?: string;
  rootUrl: string;
  redirectUris?: string[];
  webOrigins?: string[];
  enabled?: boolean;
}

export async function fn_create_keycloak_client(input: CreateKeycloakClientInput) {
  try {
    const clientConfig: ClientRepresentation = {
      clientId: input.clientId,
      name: input.name,
      description: input.description,
      rootUrl: input.rootUrl,
      baseUrl: input.rootUrl,
      enabled: input.enabled ?? true,
      protocol: 'openid-connect',
      publicClient: false,
      serviceAccountsEnabled: true,
      standardFlowEnabled: true,
      directAccessGrantsEnabled: true,
      redirectUris: input.redirectUris || [`${input.rootUrl}/*`],
      webOrigins: input.webOrigins || [input.rootUrl],
      attributes: {
        'access.token.lifespan': '300',
        'client.session.idle.timeout': '1800',
        'client.session.max.lifespan': '36000',
      },
    };

    const result = await keycloakAdmin.execute(
      async (client) => {
        return await client.clients.create(clientConfig);
      },
      'Failed to create Keycloak client'
    );

    return {
      success: true,
      clientId: result.id,
    };
  } catch (error) {
    console.error('Error creating Keycloak client:', error);
    throw error;
  }
}