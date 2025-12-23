'use server';

import { keycloakAdmin } from '@/lib/keycloak/admin-client';
import type { SessionItem } from '@/types/sessions';

export async function fn_get_client_sessions(
  clientId: string,
  clientName: string
): Promise<SessionItem[]> {
  try {
    const sessions = await keycloakAdmin.execute(
      async (client) => {
        return await client.clients.listSessions({
          id: clientId,
          first: 0,
          max: 1000,
        });
      },
      `Failed to fetch sessions for client ${clientName}`
    );

    return sessions.map((session) => ({
      id: session.id!,
      username: session.username!,
      userId: session.userId!,
      ipAddress: session.ipAddress || 'N/A',
      start: session.start!,
      lastAccess: session.lastAccess!,
      clientId: clientId,
      clientName: clientName,
    }));
  } catch (error) {
    console.error(`Error fetching sessions for client ${clientName}:`, error);
    throw error;
  }
}