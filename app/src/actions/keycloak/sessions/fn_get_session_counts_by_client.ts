'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import type { ClientSessionCount } from '@/types/sessions';

export async function fn_get_session_counts_by_client(): Promise<ClientSessionCount[]> {
  try {
    const counts: ClientSessionCount[] = [];

    // Obtener todos los clientes
    const clients = await keycloakAdmin.execute(
      async (client) => await client.clients.find(),
      'Failed to fetch clients'
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

    // Para cada cliente, obtener el conteo de sesiones
    for (const kcClient of filteredClients) {
      if (!kcClient.id) continue;

      try {
        const countResult = await keycloakAdmin.execute(
          async (client) => {
            return await client.clients.getSessionCount({ id: kcClient.id! });
          },
          `Failed to get session count for ${kcClient.clientId}`
        );

        counts.push({
          clientId: kcClient.clientId!,
          clientName: kcClient.name || kcClient.clientId!,
          sessionCount: countResult.count || 0,
        });
      } catch (error) {
        console.error(`Error getting session count for ${kcClient.clientId}:`, error);
        // Continuar con el siguiente cliente
      }
    }

    return counts.sort((a, b) => b.sessionCount - a.sessionCount);
  } catch (error) {
    console.error('Error getting session counts:', error);
    throw error;
  }
}