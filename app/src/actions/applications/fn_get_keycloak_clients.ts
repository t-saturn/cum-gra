'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';

export interface KeycloakClientSimple {
  id: string;
  client_id: string;
  name: string;
  description?: string;
  base_url?: string;
  protocol: string;
}

export async function fn_get_keycloak_clients(): Promise<KeycloakClientSimple[]> {
  try {
    const clients = await keycloakAdmin.execute(
      async (client) => await client.clients.find(),
      'Failed to fetch Keycloak clients'
    );

    // Filtrar clientes internos
    const filtered = clients.filter(
      (c) =>
        !c.clientId?.startsWith('realm-') &&
        !c.clientId?.includes('admin-cli') &&
        !c.clientId?.includes('account') &&
        !c.clientId?.includes('broker') &&
        !c.clientId?.includes('security-admin-console')
    );

    return filtered.map((c) => ({
      id: c.id!,
      client_id: c.clientId!,
      name: c.name || c.clientId!,
      description: c.description,
      base_url: c.baseUrl || c.rootUrl || '',
      protocol: c.protocol || 'openid-connect',
    }));
  } catch (error) {
    console.error('Error fetching Keycloak clients:', error);
    throw error;
  }
}