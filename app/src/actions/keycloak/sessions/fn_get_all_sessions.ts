'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import type { SessionItem } from '@/types/sessions';

export async function fn_get_all_sessions(): Promise<SessionItem[]> {
  try {
    const allSessions: SessionItem[] = [];

    // 1. Obtener todos los clientes del realm
    const clients = await keycloakAdmin.execute(
      async (client) => await client.clients.find(),
      'Failed to fetch clients'
    );

    // Filtrar clientes internos de Keycloak
    const filteredClients = clients.filter(
      (c) =>
        !c.clientId?.startsWith('realm-') &&
        !c.clientId?.includes('admin-cli') &&
        !c.clientId?.includes('account') &&
        !c.clientId?.includes('broker') &&
        !c.clientId?.includes('security-admin-console')
    );

    // 2. Para cada cliente, obtener sus sesiones
    for (const kcClient of filteredClients) {
      if (!kcClient.id) continue;

      try {
        const sessions = await keycloakAdmin.execute(
          async (client) => {
            // Obtener sesiones de usuario para este cliente
            return await client.clients.listSessions({
              id: kcClient.id!,
              first: 0,
              max: 1000, // Límite de sesiones por cliente
            });
          },
          `Failed to fetch sessions for client ${kcClient.clientId}`
        );

        // Mapear las sesiones
        sessions.forEach((session) => {
          allSessions.push({
            id: session.id!,
            username: session.username!,
            userId: session.userId!,
            ipAddress: session.ipAddress || 'N/A',
            start: session.start!,
            lastAccess: session.lastAccess!,
            clientId: kcClient.clientId!,
            clientName: kcClient.name || kcClient.clientId!,
          });
        });
      } catch (error) {
        console.error(`Error fetching sessions for client ${kcClient.clientId}:`, error);
        // Continuar con el siguiente cliente
      }
    }

    return allSessions;
  } catch (error) {
    console.error('Error fetching all sessions:', error);
    throw error;
  }
}