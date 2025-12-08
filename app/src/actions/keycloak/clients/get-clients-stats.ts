'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import { calculateClientStats } from '@/lib/keycloak/mappers';

export interface ClientStatsResult {
  success: boolean;
  data: {
    total_clients: number;
    enabled_clients: number;
    oauth_clients: number;
    saml_clients: number;
  };
  error?: string;
}

export async function getKeycloakClientsStats(): Promise<ClientStatsResult> {
  try {
    const clients = await keycloakAdmin.execute(
      async (client) => await client.clients.find(),
      'Failed to fetch Keycloak clients stats'
    );

    // Filtrar clientes internos
    const filteredClients = clients.filter(
      (c) =>
        !c.clientId?.startsWith('realm-') &&
        !c.clientId?.includes('admin-cli') &&
        !c.clientId?.includes('account') &&
        !c.clientId?.includes('broker') &&
        !c.clientId?.includes('security-admin-console')
    );

    const stats = calculateClientStats(filteredClients);

    return {
      success: true,
      data: stats,
    };
  } catch (error) {
    console.error('Error fetching Keycloak clients stats:', error);
    return {
      success: false,
      data: {
        total_clients: 0,
        enabled_clients: 0,
        oauth_clients: 0,
        saml_clients: 0,
      },
      error: error instanceof Error ? error.message : 'Unknown error',
    };
  }
}